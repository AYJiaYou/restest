package restest

import "net/http"

// Case is mainly a wrapper for http.Request
type Case struct {
	Method string
	Path   string
}

func NewCase() *Case {
	return &Case{}
}

func (c *Case) SetMethod(v string) *Case {
	c.Method = v
	return c
}

func (c *Case) SetPath(v string) *Case {
	c.Path = v
	return c
}

func (c *Case) Request() (*http.Request, error) {
	req, err := http.NewRequest(c.Method, c.Path, nil)
	return req, err
}
