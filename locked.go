package errhttp

import (
	"net/http"
)

var _ Error        = internalLocked{}
var _ ClientError  = internalLocked{}
var _ Locked       = internalLocked{}

var ErrLocked error = LockedWrap(nil)

type Locked interface {
	ClientError
	Locked()
}

type internalLocked struct {
	err error
}

func LockedWrap(err error) error {
	return internalLocked{
		err:err,
	}
}

func (receiver internalLocked) Error() string {
	return receiver.err.Error()
}

func (internalLocked) ErrHTTP() int {
	return http.StatusLocked
}

func (internalLocked) ClientError() {
	// Nothing here.
}

func (internalLocked) Locked() {
	// Nothing here.
}

func (receiver internalLocked) Unwrap() error {
	return receiver.err
}
