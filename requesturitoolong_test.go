package errhttp

import (
	"testing"
)

func TestRequestURITooLong(t *testing.T) {

	var x RequestURITooLong = internalRequestURITooLong{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
