package acars

import (
	"net/url"
)

type Request struct {
	Logon 	string
	From 	string
	To 	string
	Type 	int
	Packet  string
}


func (req *Request) ToValues() url.Values {
	valOut := url.Values{}

	valOut.Add("logon", req.Logon)
	valOut.Add("from", req.From)
	valOut.Add("type", reqTypeToString(req.Type))
	if (req.To != "") {
		valOut.Add("to", req.To)
	}
	if (req.Packet != "") {
		valOut.Add("packet", req.Packet)
	}
	return valOut
}