package run

import (
	"net/http"

	"github.com/stretchr/testify/assert"
)

// Result represnets the running-result of a Case, effectively the http response.
// You should call its Release() method after finish using it.
type Result interface {
	// Release the resources used.
	Release()
	// AssertNoError asserts there is no error in get this response.
	AssertNoError(assert.TestingT)
	// AssertCode asserts the status code of the http response.
	AssertCode(assert.TestingT, int)
}

type result struct {
	*http.Response
	err error
}

// NewResult creates an instance of the Result interface.
func NewResult(resp *http.Response, err error) Result {
	return result{resp, err}
}

func (r result) Release() {
	r.Body.Close()
}

func (r result) AssertNoError(t assert.TestingT) {
	assert.Nil(t, r.err)
}

func (r result) AssertCode(t assert.TestingT, c int) {
	assert.Equal(t, c, r.StatusCode)
}
