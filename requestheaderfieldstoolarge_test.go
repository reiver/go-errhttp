package errhttp

import (
	"testing"
)

func TestRequestHeaderFieldsTooLarge(t *testing.T) {

	var x RequestHeaderFieldsTooLarge = internalRequestHeaderFieldsTooLarge{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
