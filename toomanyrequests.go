package errhttp

import (
	"net/http"
)

var _ Error           = internalTooManyRequests{}
var _ ClientError     = internalTooManyRequests{}
var _ TooManyRequests = internalTooManyRequests{}

var ErrTooManyRequests error = TooManyRequestsWrap(nil)

type TooManyRequests interface {
	ClientError
	TooManyRequests()
}

type internalTooManyRequests struct {
	err error
}

func TooManyRequestsWrap(err error) error {
	return internalTooManyRequests{
		err:err,
	}
}

func (receiver internalTooManyRequests) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalTooManyRequests) ErrHTTP() int {
	return http.StatusTooManyRequests
}

func (internalTooManyRequests) ClientError() {
	// Nothing here.
}

func (internalTooManyRequests) TooManyRequests() {
	// Nothing here.
}

func (receiver internalTooManyRequests) Unwrap() error {
	return receiver.err
}
