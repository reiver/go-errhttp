package errhttp

import (
	"testing"
)

func TestUnauthorized(t *testing.T) {

	var x Unauthorized = internalUnauthorized{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
