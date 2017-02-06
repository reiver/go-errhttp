package errhttp

type ClientError interface {
	Error
	ClientError()
}
