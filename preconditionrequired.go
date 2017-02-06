package errhttp

type PreconditionRequired interface {
	ClientError
	PreconditionRequired()
}

type internalPreconditionRequired struct {
	err error
}

func PreconditionRequiredWrap(err error) error {
	return internalPreconditionRequired{
		err:err,
	}
}

func (receiver internalPreconditionRequired) Error() string {
	return receiver.err.Error()
}

func (receiver internalPreconditionRequired) Err() error {
	return receiver.err
}

func (internalPreconditionRequired) ClientError() {
	// Nothing here.
}

func (internalPreconditionRequired) PreconditionRequired() {
	// Nothing here.
}

