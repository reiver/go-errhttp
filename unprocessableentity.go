package errhttp

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

func (receiver internalUnprocessableEntity) Err() error {
	return receiver.err
}

func (internalUnprocessableEntity) ClientError() {
	// Nothing here.
}

func (internalUnprocessableEntity) UnprocessableEntity() {
	// Nothing here.
}
