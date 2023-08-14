package errhttp

var _ Error             = internalProxyAuthRequired{}
var _ ProxyAuthRequired = internalProxyAuthRequired{}

var ErrProxyAuthRequired error = ProxyAuthRequiredWrap(nil)

type ProxyAuthRequired interface {
	ClientError
	ProxyAuthRequired()
}

type internalProxyAuthRequired struct {
	err error
}

func ProxyAuthRequiredWrap(err error) error {
	return internalProxyAuthRequired{
		err:err,
	}
}

func (receiver internalProxyAuthRequired) Error() string {
	return receiver.err.Error()
}

func (internalProxyAuthRequired) ErrHTTP() {
	// Nothing here.
}

func (internalProxyAuthRequired) ClientError() {
	// Nothing here.
}

func (internalProxyAuthRequired) ProxyAuthRequired() {
	// Nothing here.
}

func (receiver internalProxyAuthRequired) Unwrap() error {
	return receiver.err
}
