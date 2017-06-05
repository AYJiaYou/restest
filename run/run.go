package run

// Runner run the test cases.
type Runner interface {
	Run([]TestCase) error
}
