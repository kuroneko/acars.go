package acars

import (
	//"errors"
)

// WireMsg is the interface to which all ACARS service encoded
// messages shall conform.
type WireMsg interface {
	// Convert the message to a format suitable for conveying over the wire
	WireString() string
	// Decode a message via this type
	Decode(protocolIn string) error
}
