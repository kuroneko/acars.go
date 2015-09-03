package acars

import (
	"strconv"
	"strings"
)

type CpdlcMessage struct {
	ID          int
	ReferenceID int
	RAType      string
	Body        cpdlcstring
}

const (
	CpdlcRATypeR  = "R"  // C->P Expects "ROGER"
	CpdlcRATypeWU = "WU" // C->P Expects "WILCO"/"UNABLE"
	CpdlcRATypeAN = "AN" // C->P Expects "AFFIRM"/"NEGATIVE"
	CpdlcRATypeNE = "NE" // C->P Self-closing (Not-Enabled)
	CpdlcRATypeY  = "Y"  // P->C Expects Controller Response
	CpdlcRATypeN  = "N"  // P->C Self-closing (No response expected)
)

// DecodeCpdlcMessage translates a multipart CPDLC into a cpdlc message
func DecodeCpdlcMessage(msgIn string) (cpdlc *CpdlcMessage, err error) {
	if msgIn[0] != '/' {
		return nil, ErrMalformedAcarsMessage
	}
	messageParts := strings.Split(msgIn[1:], '/')

	if msgIn[0] != "data2" {
		return nil, ErrMalformedAcarsMessage
	}
	if len(msgIn) < 5 {
		return nil, ErrMalformedAcarsMessage
	}

	cpdlc = &CpdlcMessage{}
	cpdlc.ID, err = strconv.Atoi(msgIn[1])
	if err != nil {
		return nil, ErrInvalidMessageId
	}
	cpdlc.ReferenceID, err = strconv.Atoi(msgIn[2])
	if err != nil {
		return nil, ErrInvalidMessageId
	}
	cpdlc.RAType = msgIn[3]
	cpdlc.Body = cpdlcstring(msgIn[4])

	return cpdlc
}
