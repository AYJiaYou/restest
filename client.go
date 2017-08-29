package restest

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httputil"
)

// Client is a http client with a specifi set of configurations.
type Client struct {
	hClient *http.Client
	dumpReq bool
	dumpRes bool
}

func NewClient() *Client {
	return &Client{
		hClient: &http.Client{},
		dumpReq: true,
		dumpRes: true,
	}
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

	if c.dumpReq {
		dump, err := httputil.DumpRequestOut(req, true)
		fmt.Println(">>> restest >>>>>>>>>>>>>>>>>>>>>")
		if err != nil {
			fmt.Println("Request Dump Error:", err)
		}
		fmt.Println(string(dump))
		fmt.Println("=================================")
	}

	resp, err := c.hClient.Do(req)
	if err != nil {
		return nil, err
	}

	if c.dumpRes {
		dump, err := httputil.DumpResponse(resp, true)
		fmt.Println("<<< restest <<<<<<<<<<<<<<<<<<<<<")
		if err != nil {
			fmt.Println("Request Dump Error:", err)
		}
		fmt.Println(string(dump))
		fmt.Println("=================================")
	}

	return NewResult(resp), nil
}
