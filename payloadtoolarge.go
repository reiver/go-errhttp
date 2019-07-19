package errhttp

type PayloadTooLarge interface {
	ClientError
	PayloadTooLarge()
}

type internalPayloadTooLarge struct {
	err error
}

func PayloadTooLargeWrap(err error) error {
	return internalPayloadTooLarge{
		err:err,
	}
}

func (receiver internalPayloadTooLarge) Error() string {
	return receiver.err.Error()
}

func (receiver internalPayloadTooLarge) Err() error {
	return receiver.err
}

func (internalPayloadTooLarge) ClientError() {
	// Nothing here.
}

func (internalPayloadTooLarge) PayloadTooLarge() {
	// Nothing here.
}
