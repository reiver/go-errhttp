package errhttp

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

func (receiver internalProxyAuthRequired) Err() error {
	return receiver.err
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
