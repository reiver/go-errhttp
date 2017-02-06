package errhttp

import (
	"testing"
)

func TestRequestTimeout(t *testing.T) {

	var x RequestTimeout = internalRequestTimeout{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
