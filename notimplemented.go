package errhttp

import (
	"net/http"
)

var _ Error          = internalNotImplemented{}
var _ ServerError    = internalNotImplemented{}
var _ NotImplemented = internalNotImplemented{}

var ErrNotImplemented error = NotImplementedWrap(nil)

type NotImplemented interface {
	ServerError
	NotImplemented()
}

var _ NotImplemented = internalNotImplemented{}

type internalNotImplemented struct {
	err error
}

func NotImplementedWrap(err error) error {
	return internalNotImplemented{
		err:err,
	}
}

func (receiver internalNotImplemented) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalNotImplemented) ErrHTTP() int {
	return http.StatusNotImplemented
}

func (internalNotImplemented) ServerError() {
	// Nothing here.
}

func (internalNotImplemented) NotImplemented() {
	// Nothing here.
}

func (receiver internalNotImplemented) Unwrap() error {
	return receiver.err
}
