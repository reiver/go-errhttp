package errhttp

import (
	"testing"
)

func TestLocked(t *testing.T) {

	var x Locked = internalLocked{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
