package errhttp

import (
	"testing"
)

func TestUnavailableForLegalReasons(t *testing.T) {

	var x UnavailableForLegalReasons = internalUnavailableForLegalReasons{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
