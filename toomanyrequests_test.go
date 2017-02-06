package errhttp

import (
	"testing"
)

func TestTooManyRequests(t *testing.T) {

	var x TooManyRequests = internalTooManyRequests{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
