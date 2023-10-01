package errhttp

import (
	"net/http"
)

var _ Error            = internalFailedDependency{}
var _ ClientError      = internalFailedDependency{}
var _ FailedDependency = internalFailedDependency{}

var ErrFailedDependency error = FailedDependencyWrap(nil)

type FailedDependency interface {
	ClientError
	FailedDependency()
}

type internalFailedDependency struct {
	err error
}

func FailedDependencyWrap(err error) error {
	return internalFailedDependency{
		err:err,
	}
}

func (receiver internalFailedDependency) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalFailedDependency) ErrHTTP() int {
	return http.StatusFailedDependency
}

func (internalFailedDependency) ClientError() {
	// Nothing here.
}

func (internalFailedDependency) FailedDependency() {
	// Nothing here.
}

func (receiver internalFailedDependency) Unwrap() error {
	return receiver.err
}
