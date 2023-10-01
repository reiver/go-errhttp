package errhttp

import (
	"net/http"
)

var _ Error            = internalMethodNotAllowed{}
var _ ClientError      = internalMethodNotAllowed{}
var _ MethodNotAllowed = internalMethodNotAllowed{}

var ErrMethodNotAllowed error = MethodNotAllowedWrap(nil)

type MethodNotAllowed interface {
	ClientError
	MethodNotAllowed()
}

type internalMethodNotAllowed struct {
	err error
}

func MethodNotAllowedWrap(err error) error {
	return internalMethodNotAllowed{
		err:err,
	}
}

func (receiver internalMethodNotAllowed) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalMethodNotAllowed) ErrHTTP() int {
	return http.StatusMethodNotAllowed
}

func (internalMethodNotAllowed) ClientError() {
	// Nothing here.
}

func (internalMethodNotAllowed) MethodNotAllowed() {
	// Nothing here.
}

func (receiver internalMethodNotAllowed) Unwrap() error {
	return receiver.err
}
