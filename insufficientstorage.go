package errhttp

import (
	"errors"
)

var ErrInsufficientStorage error = InsufficientStorageWrap(errors.New("Insufficient Storage"))

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
	return receiver.err.Error()
}

func (receiver internalInsufficientStorage) Err() error {
	return receiver.err
}

func (internalInsufficientStorage) ErrHTTP() {
	// Nothing here.
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
