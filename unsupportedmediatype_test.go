package errhttp

import (
	"testing"
)

func TestUnsupportedMediaType(t *testing.T) {

	var x UnsupportedMediaType = internalUnsupportedMediaType{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
