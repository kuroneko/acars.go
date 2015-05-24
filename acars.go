// Interface to Hoppie's ACARS Service
//
// Please see http://hoppie.nl/acars/ for more information

package acars

import (
	"net/http"
	"net/url"
	"io/ioutil"
)

// Server represents all of the necessary state to talk to Hoppie's ACARS
type Server struct {
	BaseUrl		string
	Logon		string
	StationName	string
}

func New(baseUrl, logon, stationName string) (srv *Server, err error) {
	srv = new(Server)

	_, err = url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	srv.BaseUrl = baseUrl
	srv.Logon = logon
	srv.StationName = stationName

	return srv, err
}

func (srv *Server) Do(req *Request) (resp string, err error) {
	hresp, err := http.PostForm(srv.BaseUrl, req.ToValues())
	if (err != nil) {
		return "", err
	}
	defer hresp.Body.Close()

	data, err := ioutil.ReadAll(hresp.Body)
	if (err != nil) {
		return "", err
	}
	return string(data), nil
}

func (srv *Server) Ping(recipient string) (extraResp []string, err error) {
	req := Request{Logon: srv.Logon, From: srv.StationName, To: recipient, Type: MsgTypePing}
	pingResp, err := srv.Do(&req)
	if err != nil {
		return nil, err
	}
	parts := CurlySplit(pingResp)
	if (parts[0] != "ok") {
		return nil, &AcarsError{errorMessage: parts[1]}
	}
	return parts[1:], nil
}

func (srv *Server) Peek() (messages []*Message, err error) {
	messages = []*Message{}
	req := Request{Logon: srv.Logon, From: srv.StationName, To: srv.StationName, Type: MsgTypePeek}
	peekMsg, err := srv.Do(&req)
	if err != nil {
		return nil, err
	}
	parts := CurlySplit(peekMsg)
	if (parts[0] != "ok") {
		return nil, &AcarsError{errorMessage: parts[1]}
	}
	for _, subMsg := range parts[1:] {
		msg := ParseMessage(subMsg)
		if msg != nil {
			messages = append(messages, msg)
		}
	}
	return messages, nil
}

func (srv *Server) Poll() (messages []*Message, err error) {
	req := Request{Logon: srv.Logon, From: srv.StationName, To: srv.StationName, Type: MsgTypePoll}
	pollMsg, err := srv.Do(&req)
	if err != nil {
		return nil, err
	}
	parts := CurlySplit(pollMsg)
	if (parts[0] != "ok") {
		return nil, &AcarsError{errorMessage: parts[1]}
	}
	for _, subMsg := range parts[1:] {
		msg := ParseMessage(subMsg)
		if msg != nil {
			messages = append(messages, msg)
		}
	}
	return messages, nil
}
