package restest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var (
	ErrNoHeader = errors.New("header not exists")
)

// Result is mainly a wrapper for http.Response
type Result struct {
	hResp *http.Response
}

func NewResult(resp *http.Response) *Result {
	if resp == nil {
		panic("response can't be nil")
	}
	return &Result{
		hResp: resp,
	}
}

func (r *Result) Release() {
	r.hResp.Body.Close()
}

func (r *Result) GetHeaderInt(h string) int {
	i, err := strconv.Atoi(r.hResp.Header.Get(h))
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return i
}

func (r *Result) GetHeaderString(h string) string {
	return r.hResp.Header.Get(h)
}

func (r *Result) GetJOSN(v interface{}) error {
	defer r.hResp.Body.Close()
	return json.NewDecoder(r.hResp.Body).Decode(v)
}

func (r *Result) GetStatus() string {
	return r.hResp.Status
}

func (r *Result) GetString() (string, error) {
	defer r.hResp.Body.Close()
	bs, err := ioutil.ReadAll(r.hResp.Body)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func (r *Result) IsSuccess() bool {
	return 200 <= r.hResp.StatusCode && r.hResp.StatusCode < 300
}

func (r *Result) Response() *http.Response {
	return r.hResp
}
