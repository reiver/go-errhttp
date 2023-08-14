package errhttp

var _ Error               = internalUnprocessableEntity{}
var _ UnprocessableEntity = internalUnprocessableEntity{}

var ErrUnprocessableEntity error = UnprocessableEntityWrap(nil)

type UnprocessableEntity interface {
	ClientError
	UnprocessableEntity()
}

type internalUnprocessableEntity struct {
	err error
}

func UnprocessableEntityWrap(err error) error {
	return internalUnprocessableEntity{
		err:err,
	}
}

func (receiver internalUnprocessableEntity) Error() string {
	return receiver.err.Error()
}

func (internalUnprocessableEntity) ErrHTTP() {
	// Nothing here.
}

func (internalUnprocessableEntity) ClientError() {
	// Nothing here.
}

func (internalUnprocessableEntity) UnprocessableEntity() {
	// Nothing here.
}

func (receiver internalUnprocessableEntity) Unwrap() error {
	return receiver.err
}
