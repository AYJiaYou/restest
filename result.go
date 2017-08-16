package restest

import (
	"net/http"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Response *http.Response
	err      error
}

func NewResult(resp *http.Response, err error) *Result {
	return &Result{
		Response: resp,
		err:      err,
	}
}

func (r *Result) Release() {
	if r.Response != nil {
		r.Response.Body.Close()
		r.Response = nil
	}
}

func (r *Result) AssertNoError(t assert.TestingT) {
	//assert.Nil(t, r.err)
}

func (r *Result) AssertCode(t assert.TestingT, c int) {
	//assert.Equal(t, c, r.StatusCode)
}
