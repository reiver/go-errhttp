package errhttp

type URITooLong interface {
	ClientError
	URITooLong()
}

type internalURITooLong struct {
	err error
}

func URITooLongWrap(err error) error {
	return internalURITooLong{
		err:err,
	}
}

func (receiver internalURITooLong) Error() string {
	return receiver.err.Error()
}

func (receiver internalURITooLong) Err() error {
	return receiver.err
}

func (internalURITooLong) ClientError() {
	// Nothing here.
}

func (internalURITooLong) URITooLong() {
	// Nothing here.
}
