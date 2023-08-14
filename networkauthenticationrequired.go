package errhttp

var _ Error                         = internalNetworkAuthenticationRequired{}
var _ NetworkAuthenticationRequired = internalNetworkAuthenticationRequired{}

var ErrNetworkAuthenticationRequired error = NetworkAuthenticationRequiredWrap(nil)

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
	return receiver.err.Error()
}

func (internalNetworkAuthenticationRequired) ErrHTTP() {
	// Nothing here.
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
