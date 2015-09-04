package acars

import (
	"testing"
)

func TestCpdlcInterfaceConformance(t *testing.T) {
	var testObj interface{} = new(CpdlcMessage)

	_, ok := testObj.(WireMsg)
	if !ok {
		t.Fatalf("CpdlcPayload does not conform to WireMsg interface")
	}
}
