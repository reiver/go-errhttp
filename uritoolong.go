package errhttp

import (
	"errors"
)

var ErrURITooLong error = URITooLongWrap(errors.New("URI Too Long"))

type URITooLong interface {
	ClientError
	URITooLong()
}

type internalURITooLong struct {
	err error
}

func URITooLongWrap(err error) error {
	return internalURITooLong{
		err:err,
	}
}

func (receiver internalURITooLong) Error() string {
	return receiver.err.Error()
}

func (receiver internalURITooLong) Err() error {
	return receiver.err
}

func (internalURITooLong) ErrHTTP() {
	// Nothing here.
}

func (internalURITooLong) ClientError() {
	// Nothing here.
}

func (internalURITooLong) URITooLong() {
	// Nothing here.
}

func (receiver internalURITooLong) Unwrap() error {
	return receiver.err
}
