package errhttp

import (
	"net/http"
)

var _ Error               = internalInternalServerError{}
var _ ServerError         = internalInternalServerError{}
var _ InternalServerError = internalInternalServerError{}

var ServerErrorInternalServerError InternalServerError = InternalServerErrorWrap(nil).(InternalServerError)
var ErrHTTPInternalServerError     Error               = ServerErrorInternalServerError
var ErrInternalServerError         error               = ServerErrorInternalServerError

type InternalServerError interface {
	ServerError
	InternalServerError()
}

var _ InternalServerError = internalInternalServerError{}

type internalInternalServerError struct {
	err error
}

func InternalServerErrorWrap(err error) error {
	return internalInternalServerError{
		err:err,
	}
}

func (receiver internalInternalServerError) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalInternalServerError) ErrHTTP() int {
	return http.StatusInternalServerError
}

func (internalInternalServerError) ServerError() {
	// Nothing here.
}

func (internalInternalServerError) InternalServerError() {
	// Nothing here.
}

func (receiver internalInternalServerError) Unwrap() error {
	return receiver.err
}
