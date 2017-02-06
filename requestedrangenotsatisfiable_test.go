package errhttp

import (
	"testing"
)

func TestRequestedRangeNotSatisfiable(t *testing.T) {

	var x RequestedRangeNotSatisfiable = internalRequestedRangeNotSatisfiable{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
