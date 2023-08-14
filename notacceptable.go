package errhttp

var _ Error         = internalNotAcceptable{}
var _ NotAcceptable = internalNotAcceptable{}

var ErrNotAcceptable error = NotAcceptableWrap(nil)

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

func (internalNotAcceptable) ErrHTTP() {
	// Nothing here.
}

func (internalNotAcceptable) ClientError() {
	// Nothing here.
}

func (internalNotAcceptable) NotAcceptable() {
	// Nothing here.
}

func (receiver internalNotAcceptable) Unwrap() error {
	return receiver.err
}
