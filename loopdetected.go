package errhttp

var _ Error        = internalLoopDetected{}
var _ LoopDetected = internalLoopDetected{}

var ErrLoopDetected error = LoopDetectedWrap(nil)

type LoopDetected interface {
	ServerError
	LoopDetected()
}

var _ LoopDetected = internalLoopDetected{}

type internalLoopDetected struct {
	err error
}

func LoopDetectedWrap(err error) error {
	return internalLoopDetected{
		err:err,
	}
}

func (receiver internalLoopDetected) Error() string {
	return receiver.err.Error()
}

func (internalLoopDetected) ErrHTTP() {
	// Nothing here.
}

func (internalLoopDetected) ServerError() {
	// Nothing here.
}

func (internalLoopDetected) LoopDetected() {
	// Nothing here.
}

func (receiver internalLoopDetected) Unwrap() error {
	return receiver.err
}
