package errhttp

type InternalServerError interface {
	ServerError
	InternalServerError()
}

type internalInternalServerError struct {
	err error
}

func InternalServerErrorWrap(err error) error {
	return internalInternalServerError{
		err:err,
	}
}

func (receiver internalInternalServerError) Error() string {
	return receiver.err.Error()
}

func (receiver internalInternalServerError) Err() error {
	return receiver.err
}

func (internalInternalServerError) ServerError() {
	// Nothing here.
}

func (internalInternalServerError) InternalServerError() {
	// Nothing here.
}
