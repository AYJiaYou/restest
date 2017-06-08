package run

import "net/http"

// Result represnets the running-result of a Case, effectively the http response.
// You should call its Release() method after finish using it.
type Result interface {
	Release()
}

type result struct {
	*http.Response
}

func (r *result) Release() {
	r.Body.Close()
}

func NewResult(resp *http.Response) Result {
	return &result{resp}
}
