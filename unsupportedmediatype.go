package errhttp

import (
	"net/http"
)

var _ Error                = internalUnsupportedMediaType{}
var _ ClientError          = internalUnsupportedMediaType{}
var _ UnsupportedMediaType = internalUnsupportedMediaType{}

var ErrUnsupportedMediaType error = UnsupportedMediaTypeWrap(nil)

type UnsupportedMediaType interface {
	ClientError
	UnsupportedMediaType()
}

type internalUnsupportedMediaType struct {
	err error
}

func UnsupportedMediaTypeWrap(err error) error {
	return internalUnsupportedMediaType{
		err:err,
	}
}

func (receiver internalUnsupportedMediaType) Error() string {
	return receiver.err.Error()
}

func (internalUnsupportedMediaType) ErrHTTP() int {
	return http.StatusUnsupportedMediaType
}

func (internalUnsupportedMediaType) ClientError() {
	// Nothing here.
}

func (internalUnsupportedMediaType) UnsupportedMediaType() {
	// Nothing here.
}

func (receiver internalUnsupportedMediaType) Unwrap() error {
	return receiver.err
}
