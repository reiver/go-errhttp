package errhttp

type NotImplemented interface {
	ServerError
	NotImplemented()
}

type internalNotImplemented struct {
	err error
}

func NotImplementedWrap(err error) error {
	return internalNotImplemented{
		err:err,
	}
}

func (receiver internalNotImplemented) Error() string {
	return receiver.err.Error()
}

func (receiver internalNotImplemented) Err() error {
	return receiver.err
}

func (internalNotImplemented) ServerError() {
	// Nothing here.
}

func (internalNotImplemented) NotImplemented() {
	// Nothing here.
}
