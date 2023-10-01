package errhttp_test

import (
	"testing"

	"net/http"

	"sourcecode.social/reiver/go-errhttp"
)

func TestErrRequestTimeout(t *testing.T) {

	var err error =   errhttp.ErrRequestTimeout
	casted       := err.(errhttp.RequestTimeout)
	var code int  =   http.StatusRequestTimeout

	{
		e := casted.Unwrap()
		if nil != e {
			t.Errorf("Expected the .Unwrap() to return nil but it actually didn't.")
			t.Logf("ERROR: (%T) %s", e, e)
			return
		}
	}

	{
		expected := code
		actual := casted.ErrHTTP()

		if expected != actual {
			t.Errorf("The actual HTTP status-code is not what was expected.")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			return
		}
	}

	{
		expected := http.StatusText(casted.ErrHTTP())
		actual := err.Error()

		if expected != actual {
			t.Errorf("The actual error message was not what was expected.")
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			return
		}
	}
}
