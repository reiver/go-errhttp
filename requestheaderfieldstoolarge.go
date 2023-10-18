package errhttp

import (
	"net/http"
)

var _ Error                       = internalRequestHeaderFieldsTooLarge{}
var _ ClientError                 = internalRequestHeaderFieldsTooLarge{}
var _ RequestHeaderFieldsTooLarge = internalRequestHeaderFieldsTooLarge{}

var ClientErrorRequestHeaderFieldsTooLarge ClientError = RequestHeaderFieldsTooLargeWrap(nil).(RequestHeaderFieldsTooLarge)
var ErrHTTPRequestHeaderFieldsTooLarge     Error       = ClientErrorRequestHeaderFieldsTooLarge
var ErrRequestHeaderFieldsTooLarge         error       = ClientErrorRequestHeaderFieldsTooLarge

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
	err := receiver.err
	if nil == err {
	return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalRequestHeaderFieldsTooLarge) ErrHTTP() int {
	return http.StatusRequestHeaderFieldsTooLarge
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
