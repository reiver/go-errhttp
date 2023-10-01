package errhttp

import (
	"net/http"
)

var _ Error         = internalNotAcceptable{}
var _ ClientError   = internalNotAcceptable{}
var _ NotAcceptable = internalNotAcceptable{}

var ErrNotAcceptable error = NotAcceptableWrap(nil)

type NotAcceptable interface {
	ClientError
	NotAcceptable()
}

type internalNotAcceptable struct {
	err error
}

func NotAcceptableWrap(err error) error {
	return internalNotAcceptable{
		err:err,
	}
}

func (receiver internalNotAcceptable) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return receiver.err.Error()
}

func (internalNotAcceptable) ErrHTTP() int {
	return http.StatusNotAcceptable
}

func (internalNotAcceptable) ClientError() {
	// Nothing here.
}

func (internalNotAcceptable) NotAcceptable() {
	// Nothing here.
}

func (receiver internalNotAcceptable) Unwrap() error {
	return receiver.err
}
