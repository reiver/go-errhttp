package errhttp

type FailedDependency interface {
	ClientError
	FailedDependency()
}

type internalFailedDependency struct {
	err error
}

func FailedDependencyWrap(err error) error {
	return internalFailedDependency{
		err:err,
	}
}

func (receiver internalFailedDependency) Error() string {
	return receiver.err.Error()
}

func (receiver internalFailedDependency) Err() error {
	return receiver.err
}

func (internalFailedDependency) ClientError() {
	// Nothing here.
}

func (internalFailedDependency) FailedDependency() {
	// Nothing here.
}

func (receiver internalFailedDependency) Unwrap() error {
	return receiver.err
}
