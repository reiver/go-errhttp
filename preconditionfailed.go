package errhttp

import (
	"net/http"
)

var _ Error              = internalPreconditionFailed{}
var _ ClientError        = internalPreconditionFailed{}
var _ PreconditionFailed = internalPreconditionFailed{}

var ClientErrorPreconditionFailed ClientError = PreconditionFailedWrap(nil).(PreconditionFailed)
var ErrHTTPPreconditionFailed     Error       = ClientErrorPreconditionFailed
var ErrPreconditionFailed         error       = ClientErrorPreconditionFailed

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
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalPreconditionFailed) ErrHTTP() int {
	return http.StatusPreconditionFailed
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
