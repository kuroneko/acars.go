package acars

import (
	"testing"
)

func TestCurlySplit(t *testing.T) {
	split := CurlySplit("ok {1 2 3} {4 5 6}")

	if len(split) != 3 {
		t.Errorf("Split into wrong number of pieces")
	}
	if split[0] != "ok" {
		t.Errorf("Unexpected first section: got \"%s\"", split[0])
	}
	if split[1] != "1 2 3" {
		t.Errorf("Unexpected second section: got \"%s\"", split[1])
	}
}

func TestNestedCurlySplit(t *testing.T) {
	split := CurlySplit("ok {1 {2 3}} {4 5 6}")

	if len(split) != 3 {
		t.Errorf("Split into wrong number of pieces")
	}
	if split[0] != "ok" {
		t.Errorf("Unexpected first section: got \"%s\"", split[0])
	}
	if split[1] != "1 {2 3}" {
		t.Fatalf("Unexpected second section: got \"%s\"", split[1])
	}

	subsplit := CurlySplit(split[1])
	if len(subsplit) != 2 {
		t.Errorf("Subsplit split into wrong number of pieces")
	}
}