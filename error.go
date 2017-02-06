package errhttp

type Error interface {
	error
	Err() error
}
