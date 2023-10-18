package errhttp

import (
	"net/http"
)

var _ Error                 = internalRequestEntityTooLarge{}
var _ ClientError           = internalRequestEntityTooLarge{}
var _ RequestEntityTooLarge = internalRequestEntityTooLarge{}

var ClientErrorRequestEntityTooLarge ClientError = RequestEntityTooLargeWrap(nil).(RequestEntityTooLarge)
var ErrHTTPRequestEntityTooLarge     Error       = ClientErrorRequestEntityTooLarge
var ErrRequestEntityTooLarge         error       = ClientErrorRequestEntityTooLarge

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
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalRequestEntityTooLarge) ErrHTTP() int {
	return http.StatusRequestEntityTooLarge
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
