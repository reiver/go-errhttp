package errhttp

import (
	"testing"
)

func TestNotAcceptable(t *testing.T) {

	var x NotAcceptable = internalNotAcceptable{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
