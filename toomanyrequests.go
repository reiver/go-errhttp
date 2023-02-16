package errhttp

import (
	"errors"
)

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

func (receiver internalTooManyRequests) Err() error {
	return receiver.err
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
