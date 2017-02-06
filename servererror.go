package errhttp

type ServerError interface {
	Error
	ServerError()
}
