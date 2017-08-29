package token

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Contexter interface {
	GetVariable(name string) (string, error)
	Calculate(alg string, params []interface{}) (string, error)
}

type ReqContexter interface {
	Contexter
	SetVariable(name string, value interface{})
	SetRequest(req *http.Request)
}

type ctxTest struct {
}

func newTestContexter() Contexter {
	return &ctxTest{}
}

func (c *ctxTest) SetVariable(name string, value interface{}) {
	panic(errors.New("not implemented"))
}

func (c *ctxTest) GetVariable(name string) (string, error) {
	debugOut("Contexter:GetVariable:", name)
	return "{" + name + "}", nil
}

func (c *ctxTest) Calculate(alg string, params []interface{}) (string, error) {
	debugOut("Contexter:Calculate:", alg, params)
	return fmt.Sprintf("{{%s:%v}}", alg, params), nil
}

type ctxReq struct {
	vars map[string]string
	req  *http.Request
}

func NewReqContexter() ReqContexter {
	return &ctxReq{}
}

func (c *ctxReq) SetVariable(name string, value interface{}) {
	if c.vars == nil {
		c.vars = make(map[string]string)
	}
	c.vars[name] = fmt.Sprintf("%v", value)
}

func (c *ctxReq) SetRequest(req *http.Request) {
	c.req = req
}

func (c *ctxReq) GetVariable(name string) (string, error) {
	switch name {
	case "ReqPath":
		if c.req == nil {
			return "", errors.New("invalid request")
		}
		return c.req.URL.Path, nil

	case "ReqBody":
		if c.req == nil {
			return "", errors.New("invalid request")
		}
		if c.req.Body == nil {
			return " ", nil
		}

		var buf bytes.Buffer
		if _, err := buf.ReadFrom(c.req.Body); err != nil {
			return "", err
		}
		if err := c.req.Body.Close(); err != nil {
			return "", err
		}
		c.req.Body = ioutil.NopCloser(&buf)
		return buf.String(), nil

	default:
		if c.vars == nil {
			return "", errors.New("unkown variable: " + name)
		}
		if v, ok := c.vars[name]; ok {
			return v, nil
		}
		return "", errors.New("unkown variable: " + name)
	}
}

func (c *ctxReq) Calculate(alg string, params []interface{}) (string, error) {
	if f, ok := _Algs[alg]; ok {
		return f(params)
	}
	return "", errors.New("unknown algorithm: " + alg)
}
