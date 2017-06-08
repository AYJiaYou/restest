package run

import "net/http"

var gClient *http.Client

// Case represents a specific http request.
type Case interface {
	Run() (Result, error)
}

type tcase struct {
	method string
	path   string
}

// NewCase creates a Case instance.
func NewCase() Case {
	return &tcase{}
}

func (c *tcase) Run() (Result, error) {
	if gClient == nil {
		gClient = &http.Client{}
	}

	req, err := http.NewRequest(c.method, c.path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := gClient.Do(req)
	if err != nil {
		return nil, err
	}
	return NewResult(resp), nil
}
