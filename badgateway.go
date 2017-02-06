package errhttp

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
	return receiver.err.Error()
}

func (receiver internalBadGateway) Err() error {
	return receiver.err
}

func (internalBadGateway) ServerError() {
	// Nothing here.
}

func (internalBadGateway) BadGateway() {
	// Nothing here.
}
