package errhttp

import (
	"testing"
)

func TestTeapot(t *testing.T) {

	var x Teapot = internalTeapot{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
