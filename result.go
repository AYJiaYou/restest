package restest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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

func (r *Result) Response() *http.Response {
	return r.hResp
}

func (r *Result) GetJOSN(v interface{}) error {
	defer r.hResp.Body.Close()
	return json.NewDecoder(r.hResp.Body).Decode(v)
}

func (r *Result) GetString() (string, error) {
	defer r.hResp.Body.Close()
	bs, err := ioutil.ReadAll(r.hResp.Body)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
