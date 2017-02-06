package errhttp

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

func (receiver internalTeapot) Err() error {
	return receiver.err
}

func (internalTeapot) ClientError() {
	// Nothing here.
}

func (internalTeapot) Teapot() {
	// Nothing here.
}
