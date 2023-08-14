package errhttp

import (
	"errors"
)

var _ Error                = internalUnsupportedMediaType{}
var _ UnsupportedMediaType = internalUnsupportedMediaType{}

var ErrUnsupportedMediaType error = UnsupportedMediaTypeWrap(errors.New("Unsupported Media Type"))

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

func (internalUnsupportedMediaType) ErrHTTP() {
	// Nothing here.
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
