package errhttp

import (
	"testing"
)

func TestFailedDependency(t *testing.T) {

	var x FailedDependency = internalFailedDependency{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
