package errhttp

import (
	"net/http"
)

var _ Error              = internalServiceUnavailable{}
var _ ServerError        = internalServiceUnavailable{}
var _ ServiceUnavailable = internalServiceUnavailable{}

var ErrServiceUnavailable error = ServiceUnavailableWrap(nil)

type ServiceUnavailable interface {
	ServerError
	ServiceUnavailable()
}

type internalServiceUnavailable struct {
	err error
}

func ServiceUnavailableWrap(err error) error {
	return internalServiceUnavailable{
		err:err,
	}
}

func (receiver internalServiceUnavailable) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalServiceUnavailable) ErrHTTP() int {
	return http.StatusServiceUnavailable
}

func (internalServiceUnavailable) ServerError() {
	// Nothing here.
}

func (internalServiceUnavailable) ServiceUnavailable() {
	// Nothing here.
}

func (receiver internalServiceUnavailable) Unwrap() error {
	return receiver.err
}
