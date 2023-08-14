package errhttp

var _ Error      = internalURITooLong{}
var _ URITooLong = internalURITooLong{}

var ErrURITooLong error = URITooLongWrap(nil)

type URITooLong interface {
	ClientError
	URITooLong()
}

type internalURITooLong struct {
	err error
}

func URITooLongWrap(err error) error {
	return internalURITooLong{
		err:err,
	}
}

func (receiver internalURITooLong) Error() string {
	return receiver.err.Error()
}

func (internalURITooLong) ErrHTTP() {
	// Nothing here.
}

func (internalURITooLong) ClientError() {
	// Nothing here.
}

func (internalURITooLong) URITooLong() {
	// Nothing here.
}

func (receiver internalURITooLong) Unwrap() error {
	return receiver.err
}
