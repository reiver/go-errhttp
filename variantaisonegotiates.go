package errhttp

type VariantAlsoNegotiates interface {
	ServerError
	VariantAlsoNegotiates()
}

type internalVariantAlsoNegotiates struct {
	err error
}

func VariantAlsoNegotiatesWrap(err error) error {
	return internalVariantAlsoNegotiates{
		err:err,
	}
}

func (receiver internalVariantAlsoNegotiates) Error() string {
	return receiver.err.Error()
}

func (receiver internalVariantAlsoNegotiates) Err() error {
	return receiver.err
}

func (internalVariantAlsoNegotiates) ServerError() {
	// Nothing here.
}

func (internalVariantAlsoNegotiates) VariantAlsoNegotiates() {
	// Nothing here.
}
