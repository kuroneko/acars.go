package acars

// Payload is the interface to which all ACARS service encoded 
// payloads shall conform.
type Payload interface {
	// the ACARS message type to be used to transmit the message
	Type() string

	// a human-readable summary of the message
	String() string

	// Convert the message to a format suitable for conveying over the wire
	WireString() string

	// Decode a message into this type
	Decode(protocolIn string) error
}

