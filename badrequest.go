package errhttp

type BadRequest interface {
	ClientError
	BadRequest()
}

type internalBadRequest struct {
	err error
}

func BadRequestWrap(err error) error {
	return internalBadRequest{
		err:err,
	}
}

func (receiver internalBadRequest) Error() string {
	return receiver.err.Error()
}

func (receiver internalBadRequest) Err() error {
	return receiver.err
}

func (internalBadRequest) ClientError() {
	// Nothing here.
}

func (internalBadRequest) BadRequest() {
	// Nothing here.
}

func (receiver internalBadRequest) Unwrap() error {
	return receiver.err
}
