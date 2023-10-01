package errhttp

import (
	"net/http"
)

var _ Error             = internalProxyAuthRequired{}
var _ ClientError       = internalProxyAuthRequired{}
var _ ProxyAuthRequired = internalProxyAuthRequired{}

var ErrProxyAuthRequired error = ProxyAuthRequiredWrap(nil)

type ProxyAuthRequired interface {
	ClientError
	ProxyAuthRequired()
}

type internalProxyAuthRequired struct {
	err error
}

func ProxyAuthRequiredWrap(err error) error {
	return internalProxyAuthRequired{
		err:err,
	}
}

func (receiver internalProxyAuthRequired) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalProxyAuthRequired) ErrHTTP() int {
	return http.StatusProxyAuthRequired
}

func (internalProxyAuthRequired) ClientError() {
	// Nothing here.
}

func (internalProxyAuthRequired) ProxyAuthRequired() {
	// Nothing here.
}

func (receiver internalProxyAuthRequired) Unwrap() error {
	return receiver.err
}
