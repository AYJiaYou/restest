package run

import "testing"

func TestRun(t *testing.T) {
	c := NewCase("GET", "http://www.example.com")
	c.AssertNoError(t)
	r := c.Run()
	defer r.Release()
	r.AssertNoError(t)
	r.AssertCode(t, 200)
}
