package errhttp

import (
	"net/http"
)

var _ Error              = internalPreconditionFailed{}
var _ ClientError        = internalPreconditionFailed{}
var _ PreconditionFailed = internalPreconditionFailed{}

var ErrPreconditionFailed error = PreconditionFailedWrap(nil)

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
