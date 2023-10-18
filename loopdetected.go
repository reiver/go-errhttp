package errhttp

import (
	"net/http"
)

var _ Error        = internalLoopDetected{}
var _ ServerError  = internalLoopDetected{}
var _ LoopDetected = internalLoopDetected{}

var ServerErrorLoopDetected ServerError = LoopDetectedWrap(nil).(LoopDetected)
var ErrHTTPLoopDetected     Error       = ServerErrorLoopDetected
var ErrLoopDetected         error       = ServerErrorLoopDetected

type LoopDetected interface {
	ServerError
	LoopDetected()
}

var _ LoopDetected = internalLoopDetected{}

type internalLoopDetected struct {
	err error
}

func LoopDetectedWrap(err error) error {
	return internalLoopDetected{
		err:err,
	}
}

func (receiver internalLoopDetected) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalLoopDetected) ErrHTTP() int {
	return http.StatusLoopDetected
}

func (internalLoopDetected) ServerError() {
	// Nothing here.
}

func (internalLoopDetected) LoopDetected() {
	// Nothing here.
}

func (receiver internalLoopDetected) Unwrap() error {
	return receiver.err
}
