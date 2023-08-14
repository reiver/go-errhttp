package errhttp

var _ Error                   = internalHTTPVersionNotSupported{}
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
	return receiver.err.Error()
}

func (internalHTTPVersionNotSupported) ErrHTTP() {
	// Nothing here.
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
