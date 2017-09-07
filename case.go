package restest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/AYJiaYou/restest/token"
)

// Case is mainly a wrapper for http.Request
type Case struct {
	Body        []byte
	ContentType string
	Host        string
	Headers     map[string]string
	Method      string
	ParamQuery  url.Values
	Path        string
	TokenHeader string
	TokenFormat string

	ctx token.ReqContexter
	tp  token.TkParser
}

// NewCase creates a Case instance.
// args: Method, Path, Host
func NewCase(args ...string) *Case {
	c := &Case{
		ctx: token.NewReqContexter(),
		tp:  token.NewParser(),
	}

	if args != nil {
		c.SetMethod(args[0])
		if len(args) > 1 {
			c.SetPath(args[1])
		}
		if len(args) > 2 {
			c.SetHost(args[2])
		}
	}

	return c
}

// Request creates http.Request instance and populates it with various information.
func (c *Case) Request() (*http.Request, error) {
	req, err := http.NewRequest(c.Method, c.getURL(), bytes.NewBuffer(c.Body))

	if c.ParamQuery != nil {
		req.URL.RawQuery = c.ParamQuery.Encode()
	}

	req.Header.Set("Content-Type", c.ContentType)
	if c.Headers != nil {
		for h, c := range c.Headers {
			req.Header.Set(h, c)
		}
	}

	if len(c.TokenHeader) > 0 {
		c.ctx.SetRequest(req)
		c.tp.SetContexter(c.ctx)
		c.tp.SetFormat(c.TokenFormat)

		tk, err := c.tp.GetToken()
		if err != nil {
			return nil, err
		}
		req.Header.Set(c.TokenHeader, tk)
	}

	return req, err
}

func (c *Case) SetBodyJSON(v interface{}) *Case {
	c.ContentType = "application/json"
	if v == nil {
		return c
	}
	bs, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	c.Body = bs
	return c
}

func (c *Case) SetHeader(header string, content interface{}) *Case {
	if c.Headers == nil {
		c.Headers = make(map[string]string)
	}
	c.Headers[header] = fmt.Sprintf("%v", content)
	return c
}

func (c *Case) SetHost(v string) *Case {
	c.Host = v
	return c
}

func (c *Case) SetMethod(v string) *Case {
	c.Method = v
	return c
}

func (c *Case) SetParamQuery(name string, value interface{}) *Case {
	if c.ParamQuery == nil {
		c.ParamQuery = url.Values{}
	}
	c.ParamQuery.Set(name, fmt.Sprintf("%v", value))
	return c
}

func (c *Case) SetPath(v string) *Case {
	c.Path = v
	return c
}

func (c *Case) SetToken(header, format string) *Case {
	c.TokenHeader, c.TokenFormat = header, format
	return c
}

// SetVariable could be used to store various info with this case
// and the stored info could be fetched with GetVariable() method.
func (c *Case) SetVariable(name string, value interface{}) *Case {
	c.ctx.SetVariable(name, value)
	return c
}

func (c *Case) getURL() string {
	if strings.HasSuffix(c.Host, "/") {
		return c.Host + c.Path
	} else {
		return c.Host + "/" + c.Path
	}
}
