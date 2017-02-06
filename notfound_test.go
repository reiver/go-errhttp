package errhttp

import (
	"testing"
)

func TestNotFound(t *testing.T) {

	var x NotFound = internalNotFound{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
