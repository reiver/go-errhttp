package errhttp

import (
	"net/http"
)

var _ Error              = internalMisdirectedRequest{}
var _ ClientError        = internalMisdirectedRequest{}
var _ MisdirectedRequest = internalMisdirectedRequest{}

var ErrMisdirectedRequest error = MisdirectedRequestWrap(nil)

type MisdirectedRequest interface {
	ClientError
	MisdirectedRequest()
}

type internalMisdirectedRequest struct {
	err error
}

func MisdirectedRequestWrap(err error) error {
	return internalMisdirectedRequest{
		err:err,
	}
}

func (receiver internalMisdirectedRequest) Error() string {
	err := receiver.err
	if nil == err {
		return http.StatusText(receiver.ErrHTTP())
	}
	return err.Error()
}

func (internalMisdirectedRequest) ErrHTTP() int {
	return http.StatusMisdirectedRequest
}

func (internalMisdirectedRequest) ClientError() {
	// Nothing here.
}

func (internalMisdirectedRequest) MisdirectedRequest() {
	// Nothing here.
}

func (receiver internalMisdirectedRequest) Unwrap() error {
	return receiver.err
}
