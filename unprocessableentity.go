package errhttp

import (
	"net/http"
)

var _ Error               = internalUnprocessableEntity{}
var _ ClientError         = internalUnprocessableEntity{}
var _ UnprocessableEntity = internalUnprocessableEntity{}

var ErrUnprocessableEntity error = UnprocessableEntityWrap(nil)

type UnprocessableEntity interface {
	ClientError
	UnprocessableEntity()
}

type internalUnprocessableEntity struct {
	err error
}

func UnprocessableEntityWrap(err error) error {
	return internalUnprocessableEntity{
		err:err,
	}
}

func (receiver internalUnprocessableEntity) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalUnprocessableEntity) ErrHTTP() int {
	return http.StatusUnprocessableEntity
}

func (internalUnprocessableEntity) ClientError() {
	// Nothing here.
}

func (internalUnprocessableEntity) UnprocessableEntity() {
	// Nothing here.
}

func (receiver internalUnprocessableEntity) Unwrap() error {
	return receiver.err
}
