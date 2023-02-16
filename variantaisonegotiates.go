package errhttp

import (
	"errors"
)

var ErrVariantAlsoNegotiates error = VariantAlsoNegotiatesWrap(errors.New("Variant Also Negotiates"))

type VariantAlsoNegotiates interface {
	ServerError
	VariantAlsoNegotiates()
}

var _ VariantAlsoNegotiates = internalVariantAlsoNegotiates{}

type internalVariantAlsoNegotiates struct {
	err error
}

func VariantAlsoNegotiatesWrap(err error) error {
	return internalVariantAlsoNegotiates{
		err:err,
	}
}

func (receiver internalVariantAlsoNegotiates) Error() string {
	return receiver.err.Error()
}

func (receiver internalVariantAlsoNegotiates) Err() error {
	return receiver.err
}

func (internalVariantAlsoNegotiates) ServerError() {
	// Nothing here.
}

func (internalVariantAlsoNegotiates) VariantAlsoNegotiates() {
	// Nothing here.
}

func (receiver internalVariantAlsoNegotiates) Unwrap() error {
	return receiver.err
}
