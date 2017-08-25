package main

import "github.com/AYJiaYou/restest/token"

func main() {
	/*client := restest.NewClient()
	client.SetHost("https://i.ayjiayou.com:8083")

	ca := restest.NewCase()
	ca.SetMethod()
	fmt.Println("hello, restest! :)")*/

	token.TestParser()
}
