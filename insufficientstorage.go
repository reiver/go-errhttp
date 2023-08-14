package errhttp

var _ Error               = internalInsufficientStorage{}
var _ ServerError         = internalInsufficientStorage{}
var _ InsufficientStorage = internalInsufficientStorage{}

var ErrInsufficientStorage error = InsufficientStorageWrap(nil)

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
