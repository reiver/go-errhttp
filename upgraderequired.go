package errhttp

import (
	"net/http"
)

var _ Error           = internalUpgradeRequired{}
var _ ClientError     = internalUpgradeRequired{}
var _ UpgradeRequired = internalUpgradeRequired{}

var ErrUpgradeRequired error = UpgradeRequiredWrap(nil)

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
	return receiver.err.Error()
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
