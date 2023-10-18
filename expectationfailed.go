package errhttp

import (
	"net/http"
)

var _ Error             = internalExpectationFailed{}
var _ ClientError       = internalExpectationFailed{}
var _ ExpectationFailed = internalExpectationFailed{}

var ClientErrorExpectationFailed ClientError = ExpectationFailedWrap(nil).(ExpectationFailed)
var ErrHTTPExpectationFailed     Error       = ClientErrorExpectationFailed
var ErrExpectationFailed         error       = ClientErrorExpectationFailed

type ExpectationFailed interface {
	ClientError
	ExpectationFailed()
}

type internalExpectationFailed struct {
	err error
}

func ExpectationFailedWrap(err error) error {
	return internalExpectationFailed{
		err:err,
	}
}

func (receiver internalExpectationFailed) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalExpectationFailed) ErrHTTP() int {
	return http.StatusExpectationFailed
}

func (internalExpectationFailed) ClientError() {
	// Nothing here.
}

func (internalExpectationFailed) ExpectationFailed() {
	// Nothing here.
}

func (receiver internalExpectationFailed) Unwrap() error {
	return receiver.err
}
