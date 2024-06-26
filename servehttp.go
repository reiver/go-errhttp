package errhttp

import (
	"net/http"
)

var _ http.Handler = ErrHTTPBadGateway
var _ http.Handler = ErrHTTPConflict
var _ http.Handler = ErrHTTPExpectationFailed
var _ http.Handler = ErrHTTPFailedDependency
var _ http.Handler = ErrHTTPForbidden
var _ http.Handler = ErrHTTPGatewayTimeout
var _ http.Handler = ErrHTTPGone
var _ http.Handler = ErrHTTPHTTPVersionNotSupported
var _ http.Handler = ErrHTTPInsufficientStorage
var _ http.Handler = ErrHTTPInternalServerError
var _ http.Handler = ErrHTTPLengthRequired
var _ http.Handler = ErrHTTPLocked
var _ http.Handler = ErrHTTPLoopDetected
var _ http.Handler = ErrHTTPMethodNotAllowed
var _ http.Handler = ErrHTTPMisdirectedRequest
//...
var _ http.Handler = ErrHTTPUpgradeRequired
var _ http.Handler = ErrHTTPVariantAlsoNegotiates

func (receiver internalBadGateway) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalBadRequest) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalConflict) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalExpectationFailed) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalFailedDependency) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalForbidden) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalGatewayTimeout) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalGone) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalHTTPVersionNotSupported) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalInsufficientStorage) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalInternalServerError) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalLengthRequired) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalLocked) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalLoopDetected) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalMethodNotAllowed) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalMisdirectedRequest) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalNetworkAuthenticationRequired) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalNotAcceptable) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalNotExtended) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalNotFound) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalNotImplemented) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalPaymentRequired) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalPreconditionFailed) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalPreconditionRequired) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalProxyAuthRequired) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalRequestedRangeNotSatisfiable) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalRequestEntityTooLarge) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalRequestHeaderFieldsTooLarge) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalRequestTimeout) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalRequestURITooLong) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalServiceUnavailable) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalTeapot) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalTooEarly) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalTooManyRequests) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalUnauthorized) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalUnavailableForLegalReasons) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalUnprocessableEntity) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalUnsupportedMediaType) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalUpgradeRequired) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}

func (receiver internalVariantAlsoNegotiates) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	http.Error(responseWriter, receiver.Error(), receiver.ErrHTTP())
}
