package errhttp

import (
	"net/http"
)

var _ Error          = internalGatewayTimeout{}
var _ ServerError    = internalGatewayTimeout{}
var _ GatewayTimeout = internalGatewayTimeout{}

var ErrGatewayTimeout error = GatewayTimeoutWrap(nil)

type GatewayTimeout interface {
	ServerError
	GatewayTimeout()
}

var _ GatewayTimeout = internalGatewayTimeout{}

type internalGatewayTimeout struct {
	err error
}

func GatewayTimeoutWrap(err error) error {
	return internalGatewayTimeout{
		err:err,
	}
}

func (receiver internalGatewayTimeout) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalGatewayTimeout) ErrHTTP() int {
	return http.StatusGatewayTimeout
}

func (internalGatewayTimeout) ServerError() {
	// Nothing here.
}

func (internalGatewayTimeout) GatewayTimeout() {
	// Nothing here.
}

func (receiver internalGatewayTimeout) Unwrap() error {
	return receiver.err
}
