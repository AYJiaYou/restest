package run

import "fmt"

// TestCase _
type TestCase interface {
	Run() (Response, error)
}

type testCase struct {
}

func (c *testCase) Run() (Response, error) {
	return nil, nil
}

func Bar() {
	fmt.Println("run.testcase.Bar()")
}
