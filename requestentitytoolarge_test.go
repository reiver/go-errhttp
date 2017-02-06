package errhttp

import (
	"testing"
)

func TestRequestEntityTooLarge(t *testing.T) {

	var x RequestEntityTooLarge = internalRequestEntityTooLarge{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
