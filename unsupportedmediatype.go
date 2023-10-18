package errhttp

import (
	"net/http"
)

var _ Error                = internalUnsupportedMediaType{}
var _ ClientError          = internalUnsupportedMediaType{}
var _ UnsupportedMediaType = internalUnsupportedMediaType{}

var ClientErrorUnsupportedMediaType ClientError = UnsupportedMediaTypeWrap(nil).(UnsupportedMediaType)
var ErrHTTPUnsupportedMediaType     Error       = ClientErrorUnsupportedMediaType
var ErrUnsupportedMediaType         error       = ClientErrorUnsupportedMediaType

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
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
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
