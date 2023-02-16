package errhttp

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

func (receiver internalNetworkAuthenticationRequired) Err() error {
	return receiver.err
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
