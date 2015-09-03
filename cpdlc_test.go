package acars

import (
	"testing"
)

func TestCpdlcInterfaceConformance(t *testing.T) {
	testObj := new(CpdlcPayload)

	_, ok := testObj.(Payload)
	if (!ok) {
		t.Fatalf("CpdlcPayload does not conform to Payload interface")
	}
}