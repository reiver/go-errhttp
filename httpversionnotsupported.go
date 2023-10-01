package errhttp

import (
	"net/http"
)

var _ Error                   = internalHTTPVersionNotSupported{}
var _ ServerError             = internalHTTPVersionNotSupported{}
var _ HTTPVersionNotSupported = internalHTTPVersionNotSupported{}

var ErrHTTPVersionNotSupported error = HTTPVersionNotSupportedWrap(nil)

type HTTPVersionNotSupported interface {
	ServerError
	HTTPVersionNotSupported()
}

var _ HTTPVersionNotSupported = internalHTTPVersionNotSupported{}

type internalHTTPVersionNotSupported struct {
	err error
}

func HTTPVersionNotSupportedWrap(err error) error {
	return internalHTTPVersionNotSupported{
		err:err,
	}
}

func (receiver internalHTTPVersionNotSupported) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalHTTPVersionNotSupported) ErrHTTP() int {
	return http.StatusHTTPVersionNotSupported
}

func (internalHTTPVersionNotSupported) ServerError() {
	// Nothing here.
}

func (internalHTTPVersionNotSupported) HTTPVersionNotSupported() {
	// Nothing here.
}

func (receiver internalHTTPVersionNotSupported) Unwrap() error {
	return receiver.err
}
