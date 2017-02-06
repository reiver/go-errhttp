package errhttp

type MethodNotAllowed interface {
	ClientError
	MethodNotAllowed()
}

type internalMethodNotAllowed struct {
	err error
}

func MethodNotAllowedWrap(err error) error {
	return internalMethodNotAllowed{
		err:err,
	}
}

func (receiver internalMethodNotAllowed) Error() string {
	return receiver.err.Error()
}

func (receiver internalMethodNotAllowed) Err() error {
	return receiver.err
}

func (internalMethodNotAllowed) ClientError() {
	// Nothing here.
}

func (internalMethodNotAllowed) MethodNotAllowed() {
	// Nothing here.
}

