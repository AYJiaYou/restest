package restest

import (
	"crypto/tls"
	"net/http"
)

// Client is a http client with a specifi set of configurations.
type Client struct {
	hClient *http.Client
}

func NewClient() *Client {
	return &Client{hClient: &http.Client{}}
}

func (c *Client) CheckServerCert(b bool) {
	if c.hClient.Transport == nil {
		c.hClient.Transport = &http.Transport{}
	}
	if trans, ok := c.hClient.Transport.(*http.Transport); ok {
		if trans.TLSClientConfig == nil {
			trans.TLSClientConfig = &tls.Config{}
		}
		trans.TLSClientConfig.InsecureSkipVerify = !b
	} else {
		panic("unkown Transport format")
	}
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
