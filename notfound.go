package errhttp

import (
	"errors"
)

var ErrNotFound error = NotFoundWrap(errors.New("Not Found"))

type NotFound  interface {
	ClientError
	NotFound ()
}

type internalNotFound  struct {
	err error
}

func NotFoundWrap(err error) error {
	return internalNotFound {
		err:err,
	}
}

func (receiver internalNotFound ) Error() string {
	return receiver.err.Error()
}

func (receiver internalNotFound ) Err() error {
	return receiver.err
}

func (internalNotFound ) ErrHTTP() {
	// Nothing here.
}

func (internalNotFound ) ClientError() {
	// Nothing here.
}

func (internalNotFound ) NotFound() {
	// Nothing here.
}

func (receiver internalNotFound) Unwrap() error {
	return receiver.err
}
