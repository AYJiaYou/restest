package token

import "fmt"

const (
	_IsDebug = true
)

func debugOut(v ...interface{}) {
	if _IsDebug {
		fmt.Println(v...)
	}
}
