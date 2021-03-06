package errhttp

type RequestHeaderFieldsTooLarge interface {
	error
	Err() error
	RequestHeaderFieldsTooLarge()
}

type internalRequestHeaderFieldsTooLarge struct {
	err error
}

func RequestHeaderFieldsTooLargeWrap(err error) error {
	return internalRequestHeaderFieldsTooLarge{
		err:err,
	}
}

func (receiver internalRequestHeaderFieldsTooLarge) Error() string {
	return receiver.err.Error()
}

func (receiver internalRequestHeaderFieldsTooLarge) Err() error {
	return receiver.err
}

func (internalRequestHeaderFieldsTooLarge) ClientError() {
	// Nothing here.
}

func (internalRequestHeaderFieldsTooLarge) RequestHeaderFieldsTooLarge() {
	// Nothing here.
}
