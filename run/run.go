package run

import "fmt"

// Runner run the test cases.
type Runner interface {
	Run([]TestCase) error
}

func Foo() {
	fmt.Println("run.foo() bbb")
}
