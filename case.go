package restest

import "net/http"

// Case is mainly a wrapper for http.Request
type Case struct {
	Method      string
	Host        string
	Path        string
	TokenHeader string
	TokenFormat string
}

func NewCase() *Case {
	return &Case{}
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

func (c *Case) getURL() string {
	return c.Host + c.Path
}

func (c *Case) Request() (*http.Request, error) {
	req, err := http.NewRequest(c.Method, c.getURL(), nil)
	return req, err
}
