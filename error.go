package errhttp

import (
	"net/http"
)

type Error interface {
	error
	http.Handler
	ErrHTTP() int
	Unwrap() error
}
