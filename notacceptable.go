package errhttp

type NotAcceptable interface {
	ClientError
	NotAcceptable()
}

type internalNotAcceptable struct {
	err error
}

func NotAcceptableWrap(err error) error {
	return internalNotAcceptable{
		err:err,
	}
}

func (receiver internalNotAcceptable) Error() string {
	return receiver.err.Error()
}

func (receiver internalNotAcceptable) Err() error {
	return receiver.err
}

func (internalNotAcceptable) ClientError() {
	// Nothing here.
}

func (internalNotAcceptable) NotAcceptable() {
	// Nothing here.
}
