package errhttp

type NotExtended interface {
	ServerError
	NotExtended()
}

type internalNotExtended struct {
	err error
}

func NotExtendedWrap(err error) error {
	return internalNotExtended{
		err:err,
	}
}

func (receiver internalNotExtended) Error() string {
	return receiver.err.Error()
}

func (receiver internalNotExtended) Err() error {
	return receiver.err
}

func (internalNotExtended) ServerError() {
	// Nothing here.
}

func (internalNotExtended) NotExtended() {
	// Nothing here.
}
