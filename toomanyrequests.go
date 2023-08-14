package errhttp

import (
	"errors"
)

var _ Error           = internalTooManyRequests{}
var _ TooManyRequests = internalTooManyRequests{}

var ErrTooManyRequests error = TooManyRequestsWrap(errors.New("Too Many Requests"))

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
	return receiver.err.Error()
}

func (internalTooManyRequests) ErrHTTP() {
	// Nothing here.
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
