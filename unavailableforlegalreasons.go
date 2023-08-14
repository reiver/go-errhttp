package errhttp

var _ Error                      = internalUnavailableForLegalReasons{}
var _ UnavailableForLegalReasons = internalUnavailableForLegalReasons{}

var ErrUnavailableForLegalReasons error = UnavailableForLegalReasonsWrap(nil)

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
	return receiver.err.Error()
}

func (internalUnavailableForLegalReasons) ErrHTTP() {
	// Nothing here.
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
