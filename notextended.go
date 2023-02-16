package errhttp

import (
	"errors"
)

var ErrNotExtended error = NotExtendedWrap(errors.New("Not Extended"))

type NotExtended interface {
	ServerError
	NotExtended()
}

var _ NotExtended = internalNotExtended{}

type internalNotExtended struct {
	err error
}

func NotExtendedWrap(err error) error {
	return internalNotExtended{
		err:err,
	}
}

func (receiver internalNotExtended) Error() string {
	return receiver.err.Error()
}

func (receiver internalNotExtended) Err() error {
	return receiver.err
}

func (internalNotExtended) ErrHTTP() {
	// Nothing here.
}

func (internalNotExtended) ServerError() {
	// Nothing here.
}

func (internalNotExtended) NotExtended() {
	// Nothing here.
}

func (receiver internalNotExtended) Unwrap() error {
	return receiver.err
}
