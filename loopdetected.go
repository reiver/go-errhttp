package errhttp

import (
	"errors"
)

var ErrLoopDetected error = LoopDetectedWrap(errors.New("Loop Detected"))

type LoopDetected interface {
	ServerError
	LoopDetected()
}

var _ LoopDetected = internalLoopDetected{}

type internalLoopDetected struct {
	err error
}

func LoopDetectedWrap(err error) error {
	return internalLoopDetected{
		err:err,
	}
}

func (receiver internalLoopDetected) Error() string {
	return receiver.err.Error()
}

func (receiver internalLoopDetected) Err() error {
	return receiver.err
}

func (internalLoopDetected) ErrHTTP() {
	// Nothing here.
}

func (internalLoopDetected) ServerError() {
	// Nothing here.
}

func (internalLoopDetected) LoopDetected() {
	// Nothing here.
}

func (receiver internalLoopDetected) Unwrap() error {
	return receiver.err
}
