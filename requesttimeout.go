package errhttp

type RequestTimeout interface {
	error
	Err() error
	RequestTimeout()
}

type internalRequestTimeout struct {
	err error
}

func RequestTimeoutWrap(err error) error {
	return internalRequestTimeout{
		err:err,
	}
}

func (receiver internalRequestTimeout) Error() string {
	return receiver.err.Error()
}

func (receiver internalRequestTimeout) Err() error {
	return receiver.err
}

func (internalRequestTimeout) ClientError() {
	// Nothing here.
}

func (internalRequestTimeout) RequestTimeout() {
	// Nothing here.
}

func (receiver internalRequestTimeout) Unwrap() error {
	return receiver.err
}
