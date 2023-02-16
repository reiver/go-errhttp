package errhttp

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

func (receiver internalHTTPVersionNotSupported) Err() error {
	return receiver.err
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
