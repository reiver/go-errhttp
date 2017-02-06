package errhttp

import (
	"testing"
)

func TestClientError(t *testing.T) {

	var x ClientError

	x = internalBadRequest{}                   // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalUnauthorized{}                 // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalPaymentRequired{}              // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalForbidden{}                    // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalNotFound{}                     // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalMethodNotAllowed{}             // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalNotAcceptable{}                // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalProxyAuthRequired{}            // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalRequestTimeout{}               // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalConflict{}                     // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalGone{}                         // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalLengthRequired{}               // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalPreconditionFailed{}           // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalRequestEntityTooLarge{}        // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalRequestURITooLong{}            // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalUnsupportedMediaType{}         // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalRequestedRangeNotSatisfiable{} // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalExpectationFailed{}            // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalTeapot{}                       // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalUnprocessableEntity{}          // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalLocked{}                       // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalFailedDependency{}             // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalUpgradeRequired{}              // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalPreconditionRequired{}         // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalTooManyRequests{}              // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalRequestHeaderFieldsTooLarge{}  // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.
	x = internalUnavailableForLegalReasons{}   // THESE ARE THE LINE THAT ACTUALLY MATTERS IN THIS TEST.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}
