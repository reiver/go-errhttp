package errhttp

import (
	"net/http"
)

var _ Error       = internalNotExtended{}
var _ ServerError = internalNotExtended{}
var _ NotExtended = internalNotExtended{}

var ErrNotExtended error = NotExtendedWrap(nil)

type NotExtended interface {
	ServerError
	NotExtended()
}

var _ NotExtended = internalNotExtended{}

type internalNotExtended struct {
	err error
}

func NotExtendedWrap(err error) error {
	return internalNotExtended{
		err:err,
	}
}

func (receiver internalNotExtended) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalNotExtended) ErrHTTP() int {
	return http.StatusNotExtended
}

func (internalNotExtended) ServerError() {
	// Nothing here.
}

func (internalNotExtended) NotExtended() {
	// Nothing here.
}

func (receiver internalNotExtended) Unwrap() error {
	return receiver.err
}
