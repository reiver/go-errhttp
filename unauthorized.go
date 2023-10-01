package errhttp

import (
	"net/http"
)

var _ Error        = internalUnauthorized{}
var _ ClientError  = internalUnauthorized{}
var _ Unauthorized = internalUnauthorized{}

var ErrUnauthorized error = UnauthorizedWrap(nil)

type Unauthorized interface {
	ClientError
	Unauthorized()
}

type internalUnauthorized struct {
	err error
}

func UnauthorizedWrap(err error) error {
	return internalUnauthorized{
		err:err,
	}
}

func (receiver internalUnauthorized) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalUnauthorized) ErrHTTP() int {
	return http.StatusUnauthorized
}

func (internalUnauthorized) ClientError() {
	// Nothing here.
}

func (internalUnauthorized) Unauthorized() {
	// Nothing here.
}

func (receiver internalUnauthorized) Unwrap() error {
	return receiver.err
}
