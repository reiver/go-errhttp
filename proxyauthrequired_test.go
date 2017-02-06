package errhttp

import (
	"testing"
)

func TestProxyAuthRequired(t *testing.T) {

	var x ProxyAuthRequired = internalProxyAuthRequired{} // THIS IS THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should not happen.")
	}
}
