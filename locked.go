package errhttp

import (
	"errors"
)

var ErrLocked error = LockedWrap(errors.New("Locked"))

type Locked interface {
	ClientError
	Locked()
}

type internalLocked struct {
	err error
}

func LockedWrap(err error) error {
	return internalLocked{
		err:err,
	}
}

func (receiver internalLocked) Error() string {
	return receiver.err.Error()
}

func (receiver internalLocked) Err() error {
	return receiver.err
}

func (internalLocked) ErrHTTP() {
	// Nothing here.
}

func (internalLocked) ClientError() {
	// Nothing here.
}

func (internalLocked) Locked() {
	// Nothing here.
}

func (receiver internalLocked) Unwrap() error {
	return receiver.err
}
