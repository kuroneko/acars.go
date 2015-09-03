package acars

import (
	"fmt"
	"github.com/kuroneko/tclmanip.go"
)

// Message represents a single subpayload from an ACARS server (as part of a peek or poll)
type Message struct {
	From    string // sending station ID
	Type    int    // message type
	Message string // payload
}

// ParseMessage takes a raw message (as recieved from the server) and parses it into our
// internal Message format
func ParseMessage(raw tclmanip.TclList) *Message {
	bits := raw.Split()
	if len(bits) < 3 {
		return nil
	}
	msg := Message{
		From:    bits[0].String(),
		Message: bits[2].String(),
	}
	msg.Type = reqTypeFromString(bits[1].String())

	return &msg
}

// String provides a human interpretable version of the CPDLC message
func (msg *Message) String() string {
	return fmt.Sprintf("%s: (%s) %s", msg.From, reqTypeToString(msg.Type), msg.Message)
}

// Decode translates the Message (payload) into a useful form if able.
func (msg *Message) Decode() (decodedMsg interface{}, err error) {
	switch msg.Type {
	case MsgTypeCPDLC:
		return DecodeCpdlcMessage(msg.Message)
	case MsgTypePosReq:
		fallthrough
	case MsgTypeTelex:
		return nil, ErrUnsupportedMessageFormat
	default:
		return nil, ErrUnknownMessageFormat
	}
}

type PollResponse struct {
	Status   string
	Messages []Message
}
