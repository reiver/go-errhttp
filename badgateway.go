package errhttp

import (
	"net/http"
)

var _ Error       = internalBadGateway{}
var _ ServerError = internalBadGateway{}
var _ BadGateway  = internalBadGateway{}

var ServerErrorBadGateway ServerError = BadGatewayWrap(nil).(BadGateway)
var ErrHTTPBadGateway     Error       = ServerErrorBadGateway
var ErrBadGateway         error       = ServerErrorBadGateway

type BadGateway interface {
	ServerError
	BadGateway()
}

type internalBadGateway struct {
	err error
}

func BadGatewayWrap(err error) error {
	return internalBadGateway{
		err:err,
	}
}

func (receiver internalBadGateway) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalBadGateway) ErrHTTP() int {
	return http.StatusBadGateway
}

func (internalBadGateway) ServerError() {
	// Nothing here.
}

func (internalBadGateway) BadGateway() {
	// Nothing here.
}

func (receiver internalBadGateway) Unwrap() error {
	return receiver.err
}
