package errhttp

var _ Error            = internalFailedDependency{}
var _ FailedDependency = internalFailedDependency{}

var ErrFailedDependency error = FailedDependencyWrap(nil)

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

func (internalFailedDependency) ErrHTTP() {
	// Nothing here.
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
