package errhttp

import (
	"net/http"
)

var _ Error               = internalInsufficientStorage{}
var _ ServerError         = internalInsufficientStorage{}
var _ InsufficientStorage = internalInsufficientStorage{}

var ServerErrorInsufficientStorage ServerError = InsufficientStorageWrap(nil).(InsufficientStorage)
var ErrHTTPInsufficientStorage     Error       = ServerErrorInsufficientStorage
var ErrInsufficientStorage         error       = ServerErrorInsufficientStorage

type InsufficientStorage interface {
	ServerError
	InsufficientStorage()
}

var _ InsufficientStorage = internalInsufficientStorage{}

type internalInsufficientStorage struct {
	err error
}

func InsufficientStorageWrap(err error) error {
	return internalInsufficientStorage{
		err:err,
	}
}

func (receiver internalInsufficientStorage) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalInsufficientStorage) ErrHTTP() int {
	return http.StatusInsufficientStorage
}

func (internalInsufficientStorage) ServerError() {
	// Nothing here.
}

func (internalInsufficientStorage) InsufficientStorage() {
	// Nothing here.
}

func (receiver internalInsufficientStorage) Unwrap() error {
	return receiver.err
}
