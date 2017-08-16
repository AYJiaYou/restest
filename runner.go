package restest

// Runner run the test cases.
type Runner interface {
	Run([]Case) error
}
