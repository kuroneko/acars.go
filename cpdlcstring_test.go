package acars

import (
	"testing"
)

func TestCpdlcStringDecode(t *testing.T) {
	thisString := CpdlcString("|@TEST@STRING@ONE")

	if thisString.String() != "/\nTEST\nSTRING\nONE" {
		t.Fatalf("Didn't decode string correctly, got: %s", thisString.String())
	}
}

func TestCpdlcStringEncode(t *testing.T) {
	thisString := NewCpdlcString("/\nTEST\nSTRING\nTWO")

	if string(thisString) != "|@TEST@STRING@TWO" {
		t.Fatalf("Didn't encode string correctly, got: %s", string(thisString))
	}
}
