package run

import (
	"net/http"

	"github.com/stretchr/testify/assert"
)

var gClient *http.Client

// Case represents a specific http request.
type Case interface {
	// Run this case, namely make the http call and get its result.
	Run() Result
	// AssertNoError asserts there is no error in creating this case.
	AssertNoError(assert.TestingT)
}

type tcase struct {
	*http.Request
	err error
}

// NewCase creates a Case instance.
func NewCase(method, url string) Case {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return tcase{nil, err}
	}
	return tcase{req, nil}
}

func (c tcase) Run() Result {
	if gClient == nil {
		gClient = &http.Client{}
	}

	resp, err := gClient.Do(c.Request)
	if err != nil {
		return NewResult(nil, err)
	}
	return NewResult(resp, nil)
}

func (c tcase) AssertNoError(t assert.TestingT) {
	assert.Nil(t, c.err)
}
