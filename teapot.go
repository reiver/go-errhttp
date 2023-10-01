package errhttp

import (
	"net/http"
)

var _ Error       = internalTeapot{}
var _ ClientError = internalTeapot{}
var _ Teapot      = internalTeapot{}

var ErrTeapot error = TeapotWrap(nil)

type Teapot interface {
	ClientError
	Teapot()
}

type internalTeapot struct {
	err error
}

func TeapotWrap(err error) error {
	return internalTeapot{
		err:err,
	}
}

func (receiver internalTeapot) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalTeapot) ErrHTTP() int {
	return http.StatusTeapot
}

func (internalTeapot) ClientError() {
	// Nothing here.
}

func (internalTeapot) Teapot() {
	// Nothing here.
}

func (receiver internalTeapot) Unwrap() error {
	return receiver.err
}
