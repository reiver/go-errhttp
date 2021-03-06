package errhttp

type UpgradeRequired interface {
	ClientError
	UpgradeRequired()
}

type internalUpgradeRequired struct {
	err error
}

func UpgradeRequiredWrap(err error) error {
	return internalUpgradeRequired{
		err:err,
	}
}

func (receiver internalUpgradeRequired) Error() string {
	return receiver.err.Error()
}

func (receiver internalUpgradeRequired) Err() error {
	return receiver.err
}

func (internalUpgradeRequired) ClientError() {
	// Nothing here.
}

func (internalUpgradeRequired) UpgradeRequired() {
	// Nothing here.
}
