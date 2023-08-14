package errhttp

var _ Error     = internalForbidden{}
var _ Forbidden = internalForbidden{}

var ErrForbidden error = ForbiddenWrap(nil)

type Forbidden interface {
	ClientError
	Forbidden()
}

type internalForbidden struct {
	err error
}

func ForbiddenWrap(err error) error {
	return internalForbidden{
		err:err,
	}
}

func (receiver internalForbidden) Error() string {
	return receiver.err.Error()
}

func (internalForbidden) ErrHTTP() {
	// Nothing here.
}

func (internalForbidden) ClientError() {
	// Nothing here.
}

func (internalForbidden) Forbidden() {
	// Nothing here.
}

func (receiver internalForbidden) Unwrap() error {
	return receiver.err
}
