package restest

import "net/http"

// Client is a http client with a specifi set of configurations.
type Client struct {
	Host string

	hClient *http.Client
}

func NewClient() *Client {
	return &Client{hClient: &http.Client{}}
}

func (c *Client) SetHost(v string) *Client {
	c.Host = v
	return c
}

func (c *Client) RunCase(ca *Case) (*Result, error) {
	req, err := ca.Request()
	if err != nil {
		return nil, err
	}
	resp, err := c.hClient.Do(req)
	if err != nil {
		return nil, err
	}
	return NewResult(resp), nil
}
