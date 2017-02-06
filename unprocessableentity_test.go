package errhttp

import (
	"testing"
)

func TestUnprocessableEntity(t *testing.T) {

	var x UnprocessableEntity = internalUnprocessableEntity{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
