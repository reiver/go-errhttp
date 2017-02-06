package errhttp

import (
	"testing"
)

func TestGone(t *testing.T) {

	var x Gone = internalGone{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
