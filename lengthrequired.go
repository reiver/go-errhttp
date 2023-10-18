package errhttp

import (
	"net/http"
)

var _ Error          = internalLengthRequired{}
var _ ClientError    = internalLengthRequired{}
var _ LengthRequired = internalLengthRequired{}

var ClientErrorLengthRequired ClientError = LengthRequiredWrap(nil).(LengthRequired)
var ErrHTTPLengthRequired     Error       = ClientErrorLengthRequired
var ErrLengthRequired         error       = ClientErrorLengthRequired

type LengthRequired interface {
	ClientError
	LengthRequired()
}

type internalLengthRequired struct {
	err error
}

func LengthRequiredWrap(err error) error {
	return internalLengthRequired{
		err:err,
	}
}

func (receiver internalLengthRequired) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalLengthRequired) ErrHTTP() int {
	return http.StatusLengthRequired
}

func (internalLengthRequired) ClientError() {
	// Nothing here.
}

func (internalLengthRequired) LengthRequired() {
	// Nothing here.
}

func (receiver internalLengthRequired) Unwrap() error {
	return receiver.err
}
