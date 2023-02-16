package errhttp

import (
	"errors"
)

var ErrNotImplemented error = NotImplementedWrap(errors.New("Not Implemented"))

type NotImplemented interface {
	ServerError
	NotImplemented()
}

var _ NotImplemented = internalNotImplemented{}

type internalNotImplemented struct {
	err error
}

func NotImplementedWrap(err error) error {
	return internalNotImplemented{
		err:err,
	}
}

func (receiver internalNotImplemented) Error() string {
	return receiver.err.Error()
}

func (receiver internalNotImplemented) Err() error {
	return receiver.err
}

func (internalNotImplemented) ServerError() {
	// Nothing here.
}

func (internalNotImplemented) NotImplemented() {
	// Nothing here.
}

func (receiver internalNotImplemented) Unwrap() error {
	return receiver.err
}
