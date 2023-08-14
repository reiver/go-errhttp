package errhttp

import (
	"errors"
)

var _ Error      = internalBadRequest{}
var _ BadRequest = internalBadRequest{}

var ErrBadRequest = BadRequestWrap(errors.New("Bad Request"))

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

func (internalBadRequest) ErrHTTP() {
	// Nothing here.
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
