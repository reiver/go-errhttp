package errhttp

type PreconditionFailed interface {
	ClientError
	PreconditionFailed()
}

type internalPreconditionFailed struct {
	err error
}

func PreconditionFailedWrap(err error) error {
	return internalPreconditionFailed{
		err:err,
	}
}

func (receiver internalPreconditionFailed) Error() string {
	return receiver.err.Error()
}

func (receiver internalPreconditionFailed) Err() error {
	return receiver.err
}

func (internalPreconditionFailed) ClientError() {
	// Nothing here.
}

func (internalPreconditionFailed) PreconditionFailed() {
	// Nothing here.
}
