package errhttp

import (
	"testing"
)

func TestPreconditionRequired(t *testing.T) {

	var x PreconditionRequired = internalPreconditionRequired{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
