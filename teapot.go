package errhttp

var _ Error  = internalTeapot{}
var _ Teapot = internalTeapot{}

var ErrTeapot error = TeapotWrap(nil)

type Teapot interface {
	ClientError
	Teapot()
}

type internalTeapot struct {
	err error
}

func TeapotWrap(err error) error {
	return internalTeapot{
		err:err,
	}
}

func (receiver internalTeapot) Error() string {
	return receiver.err.Error()
}

func (internalTeapot) ErrHTTP() {
	// Nothing here.
}

func (internalTeapot) ClientError() {
	// Nothing here.
}

func (internalTeapot) Teapot() {
	// Nothing here.
}

func (receiver internalTeapot) Unwrap() error {
	return receiver.err
}
