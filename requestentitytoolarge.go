package errhttp

type RequestEntityTooLarge interface {
	ClientError
	RequestEntityTooLarge()
}

type internalRequestEntityTooLarge struct {
	err error
}

func RequestEntityTooLargeWrap(err error) error {
	return internalRequestEntityTooLarge{
		err:err,
	}
}

func (receiver internalRequestEntityTooLarge) Error() string {
	return receiver.err.Error()
}

func (receiver internalRequestEntityTooLarge) Err() error {
	return receiver.err
}

func (internalRequestEntityTooLarge) ClientError() {
	// Nothing here.
}

func (internalRequestEntityTooLarge) RequestEntityTooLarge() {
	// Nothing here.
}

func (receiver internalRequestEntityTooLarge) Unwrap() error {
	return receiver.err
}
