package errhttp

type RequestedRangeNotSatisfiable interface {
	ClientError
	RequestedRangeNotSatisfiable()
}

type internalRequestedRangeNotSatisfiable struct {
	err error
}

func RequestedRangeNotSatisfiableWrap(err error) error {
	return internalRequestedRangeNotSatisfiable{
		err:err,
	}
}

func (receiver internalRequestedRangeNotSatisfiable) Error() string {
	return receiver.err.Error()
}

func (receiver internalRequestedRangeNotSatisfiable) Err() error {
	return receiver.err
}

func (internalRequestedRangeNotSatisfiable) ClientError() {
	// Nothing here.
}

func (internalRequestedRangeNotSatisfiable) RequestedRangeNotSatisfiable() {
	// Nothing here.
}
