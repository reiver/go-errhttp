package errhttp

import (
	"net/http"
)

var _ Error       = internalTooEarly{}
var _ ClientError = internalTooEarly{}
var _ TooEarly    = internalTooEarly{}

var ErrTooEarly = TooEarlyWrap(nil)

type TooEarly interface {
	ClientError
	TooEarly()
}

type internalTooEarly struct {
	err error
}

func TooEarlyWrap(err error) error {
	return internalTooEarly{
		err:err,
	}
}

func (receiver internalTooEarly) Error() string {
	return receiver.err.Error()
}

func (internalTooEarly) ErrHTTP() int {
	return http.StatusTooEarly
}

func (internalTooEarly) ClientError() {
	// Nothing here.
}

func (internalTooEarly) TooEarly() {
	// Nothing here.
}

func (receiver internalTooEarly) Unwrap() error {
	return receiver.err
}
