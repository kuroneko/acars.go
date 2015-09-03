package acars

import (
	"github.com/ajg/form"
	"net/url"
)

type Request struct {
	Logon  string `form:"logon"`
	From   string `form:"from"`
	To     string `form:"to"`
	Type   int    `form:"type"`
	Packet string `form:"packet"`
}

func (req *Request) ToValues() (val url.Values) {
	val, err := form.EncodeToValues(req)
	if err != nil {
		panic(err)
	}
	return val
}
