package token

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
)

type _Alg func(params []interface{}) (string, error)

var (
	_Algs = map[string]_Alg{
		"hmac_sha1":      hmac_sha1,
		"urlsafe_base64": urlsafe_base64,
	}
)

func hmac_sha1(params []interface{}) (string, error) {
	if params == nil || len(params) != 2 {
		return "", errors.New(fmt.Sprintf("invalid params for hmac_sha1: %v", params))
	}
	sk, ok := params[1].(string)
	if !ok {
		return "", errors.New(fmt.Sprintf("invalid params for hmac_sha1: %v", params))
	}
	c, ok := params[0].(string)
	if !ok {
		return "", errors.New(fmt.Sprintf("invalid params for hmac_sha1: %v", params))
	}
	hash := hmac.New(sha1.New, []byte(sk))
	hash.Write([]byte(c))
	return string(hash.Sum(nil)), nil
}

func urlsafe_base64(params []interface{}) (string, error) {
	if params == nil || len(params) != 1 {
		return "", errors.New(fmt.Sprintf("invalid params for urlsafe_base64: %v", params))
	}
	c, ok := params[0].(string)
	if !ok {
		return "", errors.New(fmt.Sprintf("invalid params for urlsafe_base64: %v", params))
	}
	return base64.URLEncoding.EncodeToString([]byte(c)), nil
}
