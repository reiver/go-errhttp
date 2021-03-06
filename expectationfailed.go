package errhttp

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

func (receiver internalExpectationFailed) Err() error {
	return receiver.err
}

func (internalExpectationFailed) ClientError() {
	// Nothing here.
}

func (internalExpectationFailed) ExpectationFailed() {
	// Nothing here.
}
