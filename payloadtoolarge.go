package errhttp

import (
	"errors"
)

var ErrPayloadTooLarge error = PayloadTooLargeWrap(errors.New("Payload Too Large"))

type PayloadTooLarge interface {
	ClientError
	PayloadTooLarge()
}

type internalPayloadTooLarge struct {
	err error
}

func PayloadTooLargeWrap(err error) error {
	return internalPayloadTooLarge{
		err:err,
	}
}

func (receiver internalPayloadTooLarge) Error() string {
	return receiver.err.Error()
}

func (receiver internalPayloadTooLarge) Err() error {
	return receiver.err
}

func (internalPayloadTooLarge) ErrHTTP() {
	// Nothing here.
}

func (internalPayloadTooLarge) ClientError() {
	// Nothing here.
}

func (internalPayloadTooLarge) PayloadTooLarge() {
	// Nothing here.
}

func (receiver internalPayloadTooLarge) Unwrap() error {
	return receiver.err
}
