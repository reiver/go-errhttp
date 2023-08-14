package errhttp

import (
	"net/http"
)

var _ Error           = internalPaymentRequired{}
var _ ClientError     = internalPaymentRequired{}
var _ PaymentRequired = internalPaymentRequired{}

var ErrPaymentRequired error = PaymentRequiredWrap(nil)

type PaymentRequired interface {
	ClientError
	PaymentRequired()
}

type internalPaymentRequired struct {
	err error
}

func PaymentRequiredWrap(err error) error {
	return internalPaymentRequired{
		err:err,
	}
}

func (receiver internalPaymentRequired) Error() string {
	return receiver.err.Error()
}

func (internalPaymentRequired) ErrHTTP() int {
	return http.StatusPaymentRequired
}


func (internalPaymentRequired) ClientError() {
	// Nothing here.
}

func (internalPaymentRequired) PaymentRequired() {
	// Nothing here.
}

func (receiver internalPaymentRequired) Unwrap() error {
	return receiver.err
}
