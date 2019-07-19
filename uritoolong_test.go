package errhttp

import (
	"testing"
)

func TestURITooLong(t *testing.T) {

	var x URITooLong = internalURITooLong{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
