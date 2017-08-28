package restest

import (
	"net/http"

	"github.com/AYJiaYou/restest/token"
)

// Case is mainly a wrapper for http.Request
type Case struct {
	Method      string
	Host        string
	Path        string
	TokenHeader string
	TokenFormat string
	ctx         token.ReqContexter
	tp          token.TkParser
}

func NewCase() *Case {
	return &Case{
		ctx: token.NewReqContexter(),
		tp:  token.NewParser(),
	}
}

func (c *Case) SetMethod(v string) *Case {
	c.Method = v
	return c
}
func (c *Case) SetHost(v string) *Case {
	c.Host = v
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
func (c *Case) SetVariable(name string, value interface{}) {
	c.ctx.SetVariable(name, value)
}

func (c *Case) getURL() string {
	return c.Host + "/" + c.Path
}

// Request creates http.Request instance and populates it with various information.
func (c *Case) Request() (*http.Request, error) {
	req, err := http.NewRequest(c.Method, c.getURL(), nil)

	// TODO: set body

	// set token
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
