package run

// Runner run the test cases.
type Runner interface {
	Run([]Case) error
}
