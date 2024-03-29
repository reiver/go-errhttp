package errhttp

import (
	"net/http"
)

var _ Error                = internalPreconditionRequired{}
var _ ClientError          = internalPreconditionRequired{}
var _ PreconditionRequired = internalPreconditionRequired{}

var ErrPreconditionRequired error = PreconditionRequiredWrap(nil)

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
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalPreconditionRequired) ErrHTTP() int {
	return http.StatusPreconditionRequired
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
