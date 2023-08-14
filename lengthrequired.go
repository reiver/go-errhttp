package errhttp

import (
	"net/http"
)

var _ Error          = internalLengthRequired{}
var _ ClientError    = internalLengthRequired{}
var _ LengthRequired = internalLengthRequired{}

var ErrLengthRequired error = LengthRequiredWrap(nil)

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
	return receiver.err.Error()
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
