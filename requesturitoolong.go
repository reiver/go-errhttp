package errhttp

import (
	"net/http"
)

var _ Error             = internalRequestURITooLong{}
var _ ClientError       = internalRequestURITooLong{}
var _ RequestURITooLong = internalRequestURITooLong{}

var ClientErrorRequestURITooLong ClientError = RequestURITooLongWrap(nil).(RequestURITooLong)
var ErrHTTPRequestURITooLong     Error       = ClientErrorRequestURITooLong
var ErrRequestURITooLong         error       = ClientErrorRequestURITooLong

type RequestURITooLong interface {
	ClientError
	RequestURITooLong()
}

type internalRequestURITooLong struct {
	err error
}

func RequestURITooLongWrap(err error) error {
	return internalRequestURITooLong{
		err:err,
	}
}

func (receiver internalRequestURITooLong) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalRequestURITooLong) ErrHTTP() int {
	return http.StatusRequestURITooLong
}

func (internalRequestURITooLong) ClientError() {
	// Nothing here.
}

func (internalRequestURITooLong) RequestURITooLong() {
	// Nothing here.
}

func (receiver internalRequestURITooLong) Unwrap() error {
	return receiver.err
}
