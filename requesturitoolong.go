package errhttp

import (
	"net/http"
)

var _ Error             = internalRequestURITooLong{}
var _ ClientError       = internalRequestURITooLong{}
var _ RequestURITooLong = internalRequestURITooLong{}

var ErrRequestURITooLong error = RequestURITooLongWrap(nil)

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
	return receiver.err.Error()
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
