package errhttp

import (
	"errors"
)

var ErrUnauthorized error = UnauthorizedWrap(errors.New("Unauthorized"))

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
	return receiver.err.Error()
}

func (receiver internalUnauthorized) Err() error {
	return receiver.err
}

func (internalUnauthorized) ErrHTTP() {
	// Nothing here.
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
