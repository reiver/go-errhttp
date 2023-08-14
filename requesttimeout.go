package errhttp

var _ Error          = internalRequestTimeout{}
var _ RequestTimeout = internalRequestTimeout{}

var ErrRequestTimeout error = RequestTimeoutWrap(nil)

type RequestTimeout interface {
	ClientError
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

func (internalRequestTimeout) ErrHTTP() {
	// Nothing here.
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
