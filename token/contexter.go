package token

import (
	"errors"
	"fmt"
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
		return "/" + c.req.RequestURI, nil

	case "ReqBody":
		if c.req == nil {
			return "", errors.New("invalid request")
		}
		if c.req.Body == nil {
			return " ", nil
		}
		return "ReqBody", nil

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
	return "", nil
}
