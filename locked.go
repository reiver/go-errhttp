package errhttp

import (
	"net/http"
)

var _ Error        = internalLocked{}
var _ ClientError  = internalLocked{}
var _ Locked       = internalLocked{}

var ClientErrorLocked ClientError = LockedWrap(nil).(Locked)
var ErrHTTPLocked     Error       = ClientErrorLocked
var ErrLocked         error       = ClientErrorLocked

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
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
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
