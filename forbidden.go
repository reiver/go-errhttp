package errhttp

import (
	"net/http"
)

var _ Error       = internalForbidden{}
var _ ClientError = internalForbidden{}
var _ Forbidden   = internalForbidden{}

var ErrForbidden error = ForbiddenWrap(nil)

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
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalForbidden) ErrHTTP() int {
	return http.StatusForbidden
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
