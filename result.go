package restest

import "net/http"

// Result is mainly a wrapper for http.Response
type Result struct {
	hResp *http.Response
}

func NewResult(resp *http.Response) *Result {
	return &Result{
		hResp: resp,
	}
}

func (r *Result) Release() {
	if r.hResp != nil {
		r.hResp.Body.Close()
		r.hResp = nil
	}
}
