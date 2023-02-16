package errhttp

import (
	"errors"
)

var ErrPreconditionFailed error = PreconditionFailedWrap(errors.New("Precondition Failed"))

type PreconditionFailed interface {
	ClientError
	PreconditionFailed()
}

type internalPreconditionFailed struct {
	err error
}

func PreconditionFailedWrap(err error) error {
	return internalPreconditionFailed{
		err:err,
	}
}

func (receiver internalPreconditionFailed) Error() string {
	return receiver.err.Error()
}

func (receiver internalPreconditionFailed) Err() error {
	return receiver.err
}

func (internalPreconditionFailed) ErrHTTP() {
	// Nothing here.
}

func (internalPreconditionFailed) ClientError() {
	// Nothing here.
}

func (internalPreconditionFailed) PreconditionFailed() {
	// Nothing here.
}

func (receiver internalPreconditionFailed) Unwrap() error {
	return receiver.err
}
