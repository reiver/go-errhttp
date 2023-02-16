package errhttp

import (
	"errors"
)

var ErrInternalServerError error = InternalServerErrorWrap(errors.New("Internal Server Error"))

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

func (receiver internalInternalServerError) Unwrap() error {
	return receiver.err
}
