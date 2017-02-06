package errhttp

import (
	"testing"
)

func TestUpgradeRequired(t *testing.T) {

	var x UpgradeRequired = internalUpgradeRequired{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
