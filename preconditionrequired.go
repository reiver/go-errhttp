package errhttp

import (
	"errors"
)

var _ Error                = internalPreconditionRequired{}
var _ PreconditionRequired = internalPreconditionRequired{}

var ErrPreconditionRequired error = PreconditionRequiredWrap(errors.New("Precondition Required"))

type PreconditionRequired interface {
	ClientError
	PreconditionRequired()
}

type internalPreconditionRequired struct {
	err error
}

func PreconditionRequiredWrap(err error) error {
	return internalPreconditionRequired{
		err:err,
	}
}

func (receiver internalPreconditionRequired) Error() string {
	return receiver.err.Error()
}

func (internalPreconditionRequired) ErrHTTP() {
	// Nothing here.
}

func (internalPreconditionRequired) ClientError() {
	// Nothing here.
}

func (internalPreconditionRequired) PreconditionRequired() {
	// Nothing here.
}

func (receiver internalPreconditionRequired) Unwrap() error {
	return receiver.err
}
