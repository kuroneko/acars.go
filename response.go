package acars

import (
	"fmt"
)

type Message struct {
	From	string
	Type 	int
	Message string
}

func ParseMessage(raw string) *Message {
	bits := CurlySplit(raw)
	if (len(bits) < 3) {
		return nil
	}
	msg := Message{
		From: bits[0],
		Message: bits[2],
	}
	msg.Type = reqTypeFromString(bits[1])

	return &msg
}

func (msg *Message) String() string {
	return fmt.Sprintf("%s: (%s) %s", msg.From, reqTypeToString(msg.Type), msg.Message)
}

type PollResponse struct {
	Status		string
	Messages	[]Message
}
