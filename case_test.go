package restest

import "testing"

func TestRun(t *testing.T) {
	c := NewCase() //"GET", "http://www.example.com")
	c.URL = "http://www.example.com"
	c.Method = "GET"
	//c.Run()
}
