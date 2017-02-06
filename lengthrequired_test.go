package errhttp

import (
	"testing"
)

func TestLengthRequired(t *testing.T) {

	var x LengthRequired = internalLengthRequired{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
