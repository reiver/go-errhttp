package errhttp

import (
	"testing"
)

func TestPayloadTooLarge(t *testing.T) {

	var x PayloadTooLarge = internalPayloadTooLarge{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
