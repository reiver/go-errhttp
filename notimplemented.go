package errhttp

var _ Error          = internalNotImplemented{}
var _ ServerError    = internalNotImplemented{}
var _ NotImplemented = internalNotImplemented{}

var ErrNotImplemented error = NotImplementedWrap(nil)

type NotImplemented interface {
	ServerError
	NotImplemented()
}

var _ NotImplemented = internalNotImplemented{}

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

func (internalNotImplemented) ErrHTTP() {
	// Nothing here.
}

func (internalNotImplemented) ServerError() {
	// Nothing here.
}

func (internalNotImplemented) NotImplemented() {
	// Nothing here.
}

func (receiver internalNotImplemented) Unwrap() error {
	return receiver.err
}
