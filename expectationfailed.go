package errhttp

var _ Error             = internalExpectationFailed{}
var _ ClientError       = internalExpectationFailed{}
var _ ExpectationFailed = internalExpectationFailed{}

var ErrExpectationFailed error = ExpectationFailedWrap(nil)

type ExpectationFailed interface {
	ClientError
	ExpectationFailed()
}

type internalExpectationFailed struct {
	err error
}

func ExpectationFailedWrap(err error) error {
	return internalExpectationFailed{
		err:err,
	}
}

func (receiver internalExpectationFailed) Error() string {
	return receiver.err.Error()
}

func (internalExpectationFailed) ErrHTTP() {
	// Nothing here.
}

func (internalExpectationFailed) ClientError() {
	// Nothing here.
}

func (internalExpectationFailed) ExpectationFailed() {
	// Nothing here.
}

func (receiver internalExpectationFailed) Unwrap() error {
	return receiver.err
}
