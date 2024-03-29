package errhttp

import (
	"net/http"
)

var _ Error                         = internalNetworkAuthenticationRequired{}
var _ ServerError                   = internalNetworkAuthenticationRequired{}
var _ NetworkAuthenticationRequired = internalNetworkAuthenticationRequired{}

var ServerErrorNetworkAuthenticationRequired ServerError = NetworkAuthenticationRequiredWrap(nil).(NetworkAuthenticationRequired)
var ErrHTTPNetworkAuthenticationRequired     Error       = ServerErrorNetworkAuthenticationRequired
var ErrNetworkAuthenticationRequired         error       = ServerErrorNetworkAuthenticationRequired

type NetworkAuthenticationRequired interface {
	ServerError
	NetworkAuthenticationRequired()
}

var _ NetworkAuthenticationRequired = internalNetworkAuthenticationRequired{}

type internalNetworkAuthenticationRequired struct {
	err error
}

func NetworkAuthenticationRequiredWrap(err error) error {
	return internalNetworkAuthenticationRequired{
		err:err,
	}
}

func (receiver internalNetworkAuthenticationRequired) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalNetworkAuthenticationRequired) ErrHTTP() int {
	return http.StatusNetworkAuthenticationRequired
}

func (internalNetworkAuthenticationRequired) ServerError() {
	// Nothing here.
}

func (internalNetworkAuthenticationRequired) NetworkAuthenticationRequired() {
	// Nothing here.
}

func (receiver internalNetworkAuthenticationRequired) Unwrap() error {
	return receiver.err
}
