package errhttp

import (
	"errors"
)

var ErrGone error = GoneWrap(errors.New("Gone"))

type Gone interface {
	ClientError()
	Gone()
}

type internalGone struct {
	err error
}

func GoneWrap(err error) error {
	return internalGone{
		err:err,
	}
}

func (receiver internalGone) Error() string {
	return receiver.err.Error()
}

func (receiver internalGone) Err() error {
	return receiver.err
}

func (internalGone) ClientError() {
	// Nothing here.
}

func (internalGone) Gone() {
	// Nothing here.
}

func (receiver internalGone) Unwrap() error {
	return receiver.err
}
