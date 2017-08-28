package token

import "fmt"

const (
	_IsDebug = false
)

func debugOut(v ...interface{}) {
	if _IsDebug {
		fmt.Println(v...)
	}
}
