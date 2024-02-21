package errhttp

import (
	"fmt"
	"net/http"
)

// Return returns the appropriate error based on the status code.
func Return(statusCode int) error {

	switch statusCode {

	case   http.StatusBadRequest:                      // 400
		return ErrBadRequest
	case   http.StatusUnauthorized:                    // 401
		return ErrUnauthorized
	case   http.StatusPaymentRequired:                 // 402
		return ErrPaymentRequired
	case   http.StatusForbidden:                       // 403
		return ErrForbidden
	case   http.StatusNotFound:                        // 404
		return ErrNotFound
	case   http.StatusMethodNotAllowed:                // 405
		return ErrMethodNotAllowed
	case   http.StatusNotAcceptable:                   // 406
		return ErrNotAcceptable
	case   http.StatusProxyAuthRequired:               // 407
		return ErrProxyAuthRequired
	case   http.StatusRequestTimeout:                  // 408
		return ErrRequestTimeout
	case   http.StatusConflict:                        // 409
		return ErrConflict
	case   http.StatusGone:                            // 410
		return ErrGone
	case   http.StatusLengthRequired:                  // 411
		return ErrLengthRequired
	case   http.StatusPreconditionFailed:              // 412
		return ErrPreconditionFailed
	case   http.StatusRequestEntityTooLarge:           // 413
		return ErrRequestEntityTooLarge
	case   http.StatusRequestURITooLong:               // 414
		return ErrRequestURITooLong
	case   http.StatusUnsupportedMediaType:            // 415
		return ErrUnsupportedMediaType
	case   http.StatusRequestedRangeNotSatisfiable:    // 416
		return ErrRequestedRangeNotSatisfiable
	case   http.StatusExpectationFailed:               // 417
		return ErrExpectationFailed
	case   http.StatusTeapot:                          // 418
		return ErrTeapot
	case   http.StatusMisdirectedRequest:              // 421
		return ErrMisdirectedRequest
	case   http.StatusUnprocessableEntity:             // 422
		return ErrUnprocessableEntity
	case   http.StatusLocked:                          // 423
		return ErrLocked
	case   http.StatusFailedDependency:                // 424
		return ErrFailedDependency
	case   http.StatusTooEarly:                        // 425
		return ErrTooEarly
	case   http.StatusUpgradeRequired:                 // 426
		return ErrUpgradeRequired
	case   http.StatusPreconditionRequired:            // 428
		return ErrPreconditionRequired
	case   http.StatusTooManyRequests:                 // 429
		return ErrTooManyRequests
	case   http.StatusRequestHeaderFieldsTooLarge:     // 431
		return ErrRequestHeaderFieldsTooLarge
	case   http.StatusUnavailableForLegalReasons:      // 451
		return ErrUnavailableForLegalReasons


	case   http.StatusInternalServerError:                // 500
		return ErrInternalServerError
	case   http.StatusNotImplemented:                     // 501
		return ErrNotImplemented
	case   http.StatusBadGateway:                         // 502
		return ErrBadGateway
	case   http.StatusServiceUnavailable:                 // 503
		return ErrServiceUnavailable
	case   http.StatusGatewayTimeout:                     // 504
		return ErrGatewayTimeout
	case   http.StatusHTTPVersionNotSupported:            // 505
		return ErrHTTPVersionNotSupported
	case   http.StatusVariantAlsoNegotiates:              // 506
		return ErrVariantAlsoNegotiates
	case   http.StatusInsufficientStorage:                // 507
		return ErrInsufficientStorage
	case   http.StatusLoopDetected:                       // 508
		return ErrLoopDetected
	case   http.StatusNotExtended:                        // 510
		return ErrNotExtended
	case   http.StatusNetworkAuthenticationRequired:      // 511
		return ErrNetworkAuthenticationRequired


	default:
		return fmt.Errorf("errhttp: unknown HTTP error — status-code %d", statusCode)
	}
}
