package errhttp

import (
	"net/http"
)

var _ Error       = internalConflict{}
var _ ClientError = internalConflict{}
var _ Conflict    = internalConflict{}

var ErrConflict = ConflictWrap(nil)

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

func (internalConflict) ErrHTTP() int {
	return http.StatusConflict
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
