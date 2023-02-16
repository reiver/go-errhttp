package errhttp

import (
	"errors"
)

var ErrServiceUnavailable error = ServiceUnavailableWrap(errors.New("Service Unavailable"))

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
	return receiver.err.Error()
}

func (receiver internalServiceUnavailable) Err() error {
	return receiver.err
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
