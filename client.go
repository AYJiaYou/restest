package restest

import (
	"crypto/tls"
	"errors"
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

func (c *Client) RunCaseFinal(ca *Case) (*Result, error) {
	rt, err := c.RunCase(ca)
	if err != nil {
		return rt, err
	}
	if !rt.IsSuccess() {
		str, err := rt.GetString()
		if err != nil {
			return rt, err
		}
		return rt, errors.New(str)
	}
	return rt, nil
}

func (c *Client) SetCheckServerCert(b bool) *Client {
	trans := c.getTransport()
	if trans.TLSClientConfig == nil {
		trans.TLSClientConfig = &tls.Config{}
	}

	trans.TLSClientConfig.InsecureSkipVerify = !b
	return c
}

func (c *Client) SetClientCert(cert tls.Certificate) *Client {
	trans := c.getTransport()
	if trans.TLSClientConfig == nil {
		trans.TLSClientConfig = &tls.Config{}
	}

	trans.TLSClientConfig.Certificates = []tls.Certificate{cert}
	//trans.TLSClientConfig.BuildNameToCertificate()
	return c
}

func (c *Client) SetClientCertFile(cert, key string) *Client {
	crt, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		panic(err)
	}
	return c.SetClientCert(crt)
}

func (c *Client) SetDumpProtocol(req, res bool) *Client {
	c.dumpReq, c.dumpRes = req, res
	return c
}

func (c *Client) getTransport() *http.Transport {
	if c.hClient.Transport == nil {
		c.hClient.Transport = &http.Transport{}
	}
	if trans, ok := c.hClient.Transport.(*http.Transport); ok {
		return trans
	}
	panic("unkown Transport format")
}
