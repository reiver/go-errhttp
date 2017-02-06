package errhttp

import (
	"testing"
)

func TestMethodNotAllowed(t *testing.T) {

	var x MethodNotAllowed = internalMethodNotAllowed{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
