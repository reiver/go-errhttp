package errhttp

import (
	"errors"
)

var _ Error    = internalConflict{}
var _ Conflict = internalConflict{}

var ErrConflict = ConflictWrap(errors.New("Conflict"))

type Conflict interface {
	ClientError
	Conflict()
}

type internalConflict struct {
	err error
}

func ConflictWrap(err error) error {
	return internalConflict{
		err:err,
	}
}

func (receiver internalConflict) Error() string {
	return receiver.err.Error()
}

func (internalConflict) ErrHTTP() {
	// Nothing here.
}

func (internalConflict) ClientError() {
	// Nothing here.
}

func (internalConflict) Conflict() {
	// Nothing here.
}

func (receiver internalConflict) Unwrap() error {
	return receiver.err
}
