package errhttp

import (
	"testing"
)

func TestExpectationFailed(t *testing.T) {

	var x ExpectationFailed = internalExpectationFailed{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
