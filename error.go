package errhttp

type Error interface {
	error
	Err() error
	Unwrap() error
}
