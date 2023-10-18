package errhttp

import (
	"net/http"
)

var _ Error                        = internalRequestedRangeNotSatisfiable{}
var _ ClientError                  = internalRequestedRangeNotSatisfiable{}
var _ RequestedRangeNotSatisfiable = internalRequestedRangeNotSatisfiable{}

var ClientErrorRequestedRangeNotSatisfiable ClientError = RequestedRangeNotSatisfiableWrap(nil).(RequestedRangeNotSatisfiable)
var ErrHTTPRequestedRangeNotSatisfiable     Error       = ClientErrorRequestedRangeNotSatisfiable
var ErrRequestedRangeNotSatisfiable         error       = ClientErrorRequestedRangeNotSatisfiable

type RequestedRangeNotSatisfiable interface {
	ClientError
	RequestedRangeNotSatisfiable()
}

type internalRequestedRangeNotSatisfiable struct {
	err error
}

func RequestedRangeNotSatisfiableWrap(err error) error {
	return internalRequestedRangeNotSatisfiable{
		err:err,
	}
}

func (receiver internalRequestedRangeNotSatisfiable) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalRequestedRangeNotSatisfiable) ErrHTTP() int {
	return http.StatusRequestedRangeNotSatisfiable
}

func (internalRequestedRangeNotSatisfiable) ClientError() {
	// Nothing here.
}

func (internalRequestedRangeNotSatisfiable) RequestedRangeNotSatisfiable() {
	// Nothing here.
}

func (receiver internalRequestedRangeNotSatisfiable) Unwrap() error {
	return receiver.err
}
