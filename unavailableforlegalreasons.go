package errhttp

import (
	"net/http"
)

var _ Error                      = internalUnavailableForLegalReasons{}
var _ ClientError                = internalUnavailableForLegalReasons{}
var _ UnavailableForLegalReasons = internalUnavailableForLegalReasons{}

var ClientErrorUnavailableForLegalReasons ClientError = UnavailableForLegalReasonsWrap(nil).(UnavailableForLegalReasons)
var ErrHTTPUnavailableForLegalReasons     Error       = ClientErrorUnavailableForLegalReasons
var ErrUnavailableForLegalReasons         error       = ClientErrorUnavailableForLegalReasons

type UnavailableForLegalReasons interface {
	ClientError
	UnavailableForLegalReasons()
}

type internalUnavailableForLegalReasons struct {
	err error
}

func UnavailableForLegalReasonsWrap(err error) error {
	return internalUnavailableForLegalReasons{
		err:err,
	}
}

func (receiver internalUnavailableForLegalReasons) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalUnavailableForLegalReasons) ErrHTTP() int {
	return http.StatusUnavailableForLegalReasons
}

func (internalUnavailableForLegalReasons) ClientError() {
	// Nothing here.
}

func (internalUnavailableForLegalReasons) UnavailableForLegalReasons() {
	// Nothing here.
}

func (receiver internalUnavailableForLegalReasons) Unwrap() error {
	return receiver.err
}
