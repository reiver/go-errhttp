package errhttp

var _ Error                 = internalVariantAlsoNegotiates{}
var _ ServerError           = internalVariantAlsoNegotiates{}
var _ VariantAlsoNegotiates = internalVariantAlsoNegotiates{}

var ErrVariantAlsoNegotiates error = VariantAlsoNegotiatesWrap(nil)

type VariantAlsoNegotiates interface {
	ServerError
	VariantAlsoNegotiates()
}

var _ VariantAlsoNegotiates = internalVariantAlsoNegotiates{}

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

func (internalVariantAlsoNegotiates) ErrHTTP() {
	// Nothing here.
}

func (internalVariantAlsoNegotiates) ServerError() {
	// Nothing here.
}

func (internalVariantAlsoNegotiates) VariantAlsoNegotiates() {
	// Nothing here.
}

func (receiver internalVariantAlsoNegotiates) Unwrap() error {
	return receiver.err
}
