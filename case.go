package restest

import "net/http"

var gClient *http.Client

type Case struct {
	URL    string
	Method string
}

func NewCase() *Case {
	return &Case{}
}

func (c *Case) Run() *Result {
	return NewResult(runCase(c))
}
