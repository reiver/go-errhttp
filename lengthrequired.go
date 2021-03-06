package errhttp

type LengthRequired interface {
	ClientError
	LengthRequired()
}

type internalLengthRequired struct {
	err error
}

func LengthRequiredWrap(err error) error {
	return internalLengthRequired{
		err:err,
	}
}

func (receiver internalLengthRequired) Error() string {
	return receiver.err.Error()
}

func (receiver internalLengthRequired) Err() error {
	return receiver.err
}

func (internalLengthRequired) ClientError() {
	// Nothing here.
}

func (internalLengthRequired) LengthRequired() {
	// Nothing here.
}

