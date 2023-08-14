package errhttp

var _ Error                       = internalRequestHeaderFieldsTooLarge{}
var _ RequestHeaderFieldsTooLarge = internalRequestHeaderFieldsTooLarge{}

var ErrRequestHeaderFieldsTooLarge error = RequestHeaderFieldsTooLargeWrap(nil)

type RequestHeaderFieldsTooLarge interface {
	ClientError
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

func (internalRequestHeaderFieldsTooLarge) ErrHTTP() {
	// Nothing here.
}

func (internalRequestHeaderFieldsTooLarge) ClientError() {
	// Nothing here.
}

func (internalRequestHeaderFieldsTooLarge) RequestHeaderFieldsTooLarge() {
	// Nothing here.
}

func (receiver internalRequestHeaderFieldsTooLarge) Unwrap() error {
	return receiver.err
}
