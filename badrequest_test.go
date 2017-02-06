package errhttp

import (
	"testing"
)

func TestBadRequest(t *testing.T) {

	var x BadRequest = internalBadRequest{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
