package errhttp

import (
	"testing"
)

func TestConflict(t *testing.T) {

	var x Conflict = internalConflict{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
