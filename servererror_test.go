package errhttp

import (
	"testing"
)

func TestServerError(t *testing.T) {

	var x ServerError

	x = internalInternalServerError{}           // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalNotImplemented{}                // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalBadGateway{}                    // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalServiceUnavailable{}            // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalGatewayTimeout{}                // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalHTTPVersionNotSupported{}       // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalVariantAlsoNegotiates{}         // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalInsufficientStorage{}           // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalLoopDetected{}                  // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalNotExtended{}                   // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalNetworkAuthenticationRequired{} // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}
