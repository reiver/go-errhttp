package errhttp

import (
	"errors"
)

var ErrForbidden error = ForbiddenWrap(errors.New("Forbidden"))

type Forbidden interface {
	ClientError
	Forbidden()
}

type internalForbidden struct {
	err error
}

func ForbiddenWrap(err error) error {
	return internalForbidden{
		err:err,
	}
}

func (receiver internalForbidden) Error() string {
	return receiver.err.Error()
}

func (receiver internalForbidden) Err() error {
	return receiver.err
}

func (internalForbidden) ClientError() {
	// Nothing here.
}

func (internalForbidden) Forbidden() {
	// Nothing here.
}

func (receiver internalForbidden) Unwrap() error {
	return receiver.err
}
