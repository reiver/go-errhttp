package errhttp

import (
	"errors"
)

var _ Error            = internalMethodNotAllowed{}
var _ MethodNotAllowed = internalMethodNotAllowed{}

var ErrMethodNotAllowed error = MethodNotAllowedWrap(errors.New("Method Not Allowed"))

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
	return receiver.err.Error()
}

func (internalMethodNotAllowed) ErrHTTP() {
	// Nothing here.
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
