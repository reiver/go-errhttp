package errhttp

type Error interface {
	error
	ErrHTTP()
	Unwrap() error
}
