package errhttp

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
	return receiver.err.Error()
}

func (receiver internalGatewayTimeout) Err() error {
	return receiver.err
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
