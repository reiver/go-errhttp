package errhttp

import (
	"errors"
)

var ErrPaymentRequired error = PaymentRequiredWrap(errors.New("Payment Required"))

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

func (receiver internalPaymentRequired) Err() error {
	return receiver.err
}

func (internalPaymentRequired) ErrHTTP() {
	// Nothing here.
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
