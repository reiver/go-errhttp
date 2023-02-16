package errhttp

import (
	"errors"
)

var ErrLengthRequired error = LengthRequiredWrap(errors.New("Length Required"))

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

func (receiver internalLengthRequired) Err() error {
	return receiver.err
}

func (internalLengthRequired) ErrHTTP() {
	// Nothing here.
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
