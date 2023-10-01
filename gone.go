package errhttp

import (
	"net/http"
)

var _ Error       = internalGone{}
var _ ClientError = internalGone{}
var _ Gone        = internalGone{}

var ErrGone error = GoneWrap(nil)

type Gone interface {
	ClientError()
	Gone()
}

type internalGone struct {
	err error
}

func GoneWrap(err error) error {
	return internalGone{
		err:err,
	}
}

func (receiver internalGone) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalGone) ErrHTTP() int {
	return http.StatusGone
}

func (internalGone) ClientError() {
	// Nothing here.
}

func (internalGone) Gone() {
	// Nothing here.
}

func (receiver internalGone) Unwrap() error {
	return receiver.err
}
