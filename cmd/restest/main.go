package main

import (
	"fmt"

	"github.com/AYJiaYou/restest"
)

func main() {
	client := restest.NewClient()
	client.SetHost("https://i.ayjiayou.com:8083")

	ca := restest.NewCase()
	ca.SetMethod()
	fmt.Println("hello, restest! :)")

}
