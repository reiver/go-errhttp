package errhttp

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

func (receiver internalUnavailableForLegalReasons) Err() error {
	return receiver.err
}

func (internalUnavailableForLegalReasons) ClientError() {
	// Nothing here.
}

func (internalUnavailableForLegalReasons) UnavailableForLegalReasons() {
	// Nothing here.
}

