package acars

import (
	"strconv"
	"strings"
)

type CpdlcMessage struct {
	ID          int
	ReferenceID *int
	RAType      string
	Body        CpdlcString
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
	cpdlc = &CpdlcMessage{}
	err = cpdlc.Decode(msgIn)
	if err != nil {
		// don't return incomplete/corrupt messages
		cpdlc = nil
	}
	return cpdlc, err
}

func (cpdlc *CpdlcMessage) Decode(msgIn string) (err error) {
	if msgIn[0] != '/' {
		return ErrMalformedAcarsMessage
	}
	messageParts := strings.Split(msgIn[1:], "/")

	if messageParts[0] != "data2" {
		return ErrMalformedAcarsMessage
	}
	if len(messageParts) < 5 {
		return ErrMalformedAcarsMessage
	}

	cpdlc.ID, err = strconv.Atoi(messageParts[1])
	if err != nil {
		return ErrInvalidMessageId
	}
	if len(messageParts[2]) > 0 {
		msgRef, err := strconv.Atoi(messageParts[2])
		if err != nil {
			return ErrInvalidMessageId
		}
		cpdlc.ReferenceID = &msgRef
	} else {
		cpdlc.ReferenceID = nil
	}
	cpdlc.RAType = messageParts[3]
	cpdlc.Body = CpdlcString(messageParts[4])

	return nil
}

func (cpdlc *CpdlcMessage) WireString() string {
	rvlist := make([]string, 5)

	rvlist[0] = "data2"
	rvlist[1] = strconv.Itoa(cpdlc.ID)
	if cpdlc.ReferenceID != nil {
		rvlist[2] = strconv.Itoa(*cpdlc.ReferenceID)
	}
	rvlist[3] = cpdlc.RAType
	rvlist[4] = string(cpdlc.Body)

	return "/" + strings.Join(rvlist, "/")
}
