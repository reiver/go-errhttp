package errhttp

type Error interface {
	error
	ErrHTTP() int
	Unwrap() error
}
