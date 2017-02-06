package errhttp

import (
	"testing"
)

func TestForbidden(t *testing.T) {

	var x Forbidden = internalForbidden{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
