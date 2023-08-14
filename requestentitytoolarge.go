package errhttp

import (
	"errors"
)

var _ Error                 = internalRequestEntityTooLarge{}
var _ RequestEntityTooLarge = internalRequestEntityTooLarge{}

var ErrRequestEntityTooLarge error = RequestEntityTooLargeWrap(errors.New("Request Entity Too Large"))

type RequestEntityTooLarge interface {
	ClientError
	RequestEntityTooLarge()
}

type internalRequestEntityTooLarge struct {
	err error
}

func RequestEntityTooLargeWrap(err error) error {
	return internalRequestEntityTooLarge{
		err:err,
	}
}

func (receiver internalRequestEntityTooLarge) Error() string {
	return receiver.err.Error()
}

func (internalRequestEntityTooLarge) ErrHTTP() {
	// Nothing here.
}

func (internalRequestEntityTooLarge) ClientError() {
	// Nothing here.
}

func (internalRequestEntityTooLarge) RequestEntityTooLarge() {
	// Nothing here.
}

func (receiver internalRequestEntityTooLarge) Unwrap() error {
	return receiver.err
}
