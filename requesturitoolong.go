package errhttp

import (
	"errors"
)

var _ Error             = internalRequestURITooLong{}
var _ RequestURITooLong = internalRequestURITooLong{}

var ErrRequestURITooLong error = RequestURITooLongWrap(errors.New("Request URI Too Long"))

type RequestURITooLong interface {
	ClientError
	RequestURITooLong()
}

type internalRequestURITooLong struct {
	err error
}

func RequestURITooLongWrap(err error) error {
	return internalRequestURITooLong{
		err:err,
	}
}

func (receiver internalRequestURITooLong) Error() string {
	return receiver.err.Error()
}

func (internalRequestURITooLong) ErrHTTP() {
	// Nothing here.
}

func (internalRequestURITooLong) ClientError() {
	// Nothing here.
}

func (internalRequestURITooLong) RequestURITooLong() {
	// Nothing here.
}

func (receiver internalRequestURITooLong) Unwrap() error {
	return receiver.err
}
