package errhttp

import (
	"net/http"
)

var _ Error       = internalBadRequest{}
var _ ClientError = internalBadRequest{}
var _ BadRequest  = internalBadRequest{}

var ClientErrorBadRequest ClientError = BadRequestWrap(nil).(BadRequest)
var ErrHTTPBadRequest     Error       = ClientErrorBadRequest
var ErrBadRequest         error       = ClientErrorBadRequest

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
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalBadRequest) ErrHTTP() int {
	return http.StatusBadRequest
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
