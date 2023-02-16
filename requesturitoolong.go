package errhttp

import (
	"errors"
)

var ErrRequestURITooLong error = RequestURITooLongWrap(errors.New("Request URI Too Long"))

type RequestURITooLong interface {
	error
	Err() error
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

func (receiver internalRequestURITooLong) Err() error {
	return receiver.err
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
