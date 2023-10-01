package errhttp

import (
	"net/http"
)

var _ Error       = internalConflict{}
var _ ClientError = internalConflict{}
var _ Conflict    = internalConflict{}

var ErrConflict error = ConflictWrap(nil)

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
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
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
