package errhttp

import (
	"net/http"
)

var _ Error          = internalRequestTimeout{}
var _ ClientError    = internalRequestTimeout{}
var _ RequestTimeout = internalRequestTimeout{}

var ErrRequestTimeout error = RequestTimeoutWrap(nil)

type RequestTimeout interface {
	ClientError
	RequestTimeout()
}

type internalRequestTimeout struct {
	err error
}

func RequestTimeoutWrap(err error) error {
	return internalRequestTimeout{
		err:err,
	}
}

func (receiver internalRequestTimeout) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalRequestTimeout) ErrHTTP() int {
	return http.StatusRequestTimeout
}

func (internalRequestTimeout) ClientError() {
	// Nothing here.
}

func (internalRequestTimeout) RequestTimeout() {
	// Nothing here.
}

func (receiver internalRequestTimeout) Unwrap() error {
	return receiver.err
}
