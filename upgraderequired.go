package errhttp

import (
	"net/http"
)

var _ Error           = internalUpgradeRequired{}
var _ ClientError     = internalUpgradeRequired{}
var _ UpgradeRequired = internalUpgradeRequired{}

var ClientErrorUpgradeRequired ClientError = UpgradeRequiredWrap(nil).(UpgradeRequired)
var ErrHTTPUpgradeRequired     Error       = ClientErrorUpgradeRequired
var ErrUpgradeRequired         error       = ClientErrorUpgradeRequired

type UpgradeRequired interface {
	ClientError
	UpgradeRequired()
}

type internalUpgradeRequired struct {
	err error
}

func UpgradeRequiredWrap(err error) error {
	return internalUpgradeRequired{
		err:err,
	}
}

func (receiver internalUpgradeRequired) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalUpgradeRequired) ErrHTTP() int {
	return http.StatusUpgradeRequired
}

func (internalUpgradeRequired) ClientError() {
	// Nothing here.
}

func (internalUpgradeRequired) UpgradeRequired() {
	// Nothing here.
}

func (receiver internalUpgradeRequired) Unwrap() error {
	return receiver.err
}
