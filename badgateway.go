package errhttp

import (
	"errors"
)

var _ Error      = internalBadGateway{}
var _ BadGateway = internalBadGateway{}

var ErrBadGateway error = BadGatewayWrap(errors.New("Bad Gateway"))

type BadGateway interface {
	ServerError
	BadGateway()
}

var _ BadGateway = internalBadGateway{}

type internalBadGateway struct {
	err error
}

func BadGatewayWrap(err error) error {
	return internalBadGateway{
		err:err,
	}
}

func (receiver internalBadGateway) Error() string {
	return receiver.err.Error()
}

func (internalBadGateway) ErrHTTP() {
	// Nothing here.
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
