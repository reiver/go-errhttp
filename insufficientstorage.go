package errhttp

type InsufficientStorage interface {
	ServerError
	InsufficientStorage()
}

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

func (internalInsufficientStorage) ServerError() {
	// Nothing here.
}

func (internalInsufficientStorage) InsufficientStorage() {
	// Nothing here.
}
