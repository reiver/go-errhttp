package errhttp

import (
	"errors"
)

var _ Error                        = internalRequestedRangeNotSatisfiable{}
var _ RequestedRangeNotSatisfiable = internalRequestedRangeNotSatisfiable{}

var ErrRequestedRangeNotSatisfiable error = RequestedRangeNotSatisfiableWrap(errors.New("Requested Range Not Satisfiable"))

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

func (internalRequestedRangeNotSatisfiable) ErrHTTP() {
	// Nothing here.
}

func (internalRequestedRangeNotSatisfiable) ClientError() {
	// Nothing here.
}

func (internalRequestedRangeNotSatisfiable) RequestedRangeNotSatisfiable() {
	// Nothing here.
}

func (receiver internalRequestedRangeNotSatisfiable) Unwrap() error {
	return receiver.err
}
