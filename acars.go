// Interface to Hoppie's ACARS Service
//
// Please see http://hoppie.nl/acars/ for more information

package acars

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	ErrMalformedAcarsMessage    = errors.New("Malformed ACARS message recieved")
	ErrInvalidMessageId         = errors.New("Message ID Invalid")
	ErrUnknownMessageFormat     = errors.New("Unknown Message Format")
	ErrUnsupportedMessageFormat = errors.New("Unsupported Message Format")
)

// Server represents all of the necessary state to talk to Hoppie's ACARS
type Server struct {
	BaseUrl     string
	Logon       string
	StationName string
}

// New creates a new Server with which you can communicate with an ACARS server.
//
// It will validate the URL provided in baseUrl.  It currently does not validate the logon.
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

// Do performs an ACARS operation.
//
// req contains the fully populated ACARS request (logon, from, to, type, packet)
// resp contains the data sent back by the server encapsulated as a tclmanip.TclList.
// err contains the error if one was encountered during the request.
func (srv *Server) Do(req *Request) (resp tclmanip.TclList, err error) {
	hresp, err := http.PostForm(srv.BaseUrl, req.ToValues())
	if err != nil {
		return "", err
	}
	defer hresp.Body.Close()

	data, err := ioutil.ReadAll(hresp.Body)
	if err != nil {
		return "", err
	}
	return tclmanip.TclList(data), nil
}

// Ping performs an ACARS ping to the recipient specified.
//
// extraResp contains the response after the "ok", if recieved.
func (srv *Server) Ping(recipient string) (extraResp tclmanip.TclList, err error) {
	req := Request{Logon: srv.Logon, From: srv.StationName, To: recipient, Type: MsgTypePing}
	pingResp, err := srv.Do(&req)
	if err != nil {
		return nil, err
	}
	pingRespParts := pingResp.Split()
	if pingRespParts[0] != "ok" {
		return nil, &AcarsError{errorMessage: pingRespParts[1]}
	}
	return tclmanip.Join(pingRespParts[1:]), nil
}

// Peek performs an ACARS Peek (fetch without update)
//
// messages contains the list of responses provided by the server
func (srv *Server) Peek() (messages []*Message, err error) {
	messages = []*Message{}
	req := Request{Logon: srv.Logon, From: srv.StationName, To: srv.StationName, Type: MsgTypePeek}
	peekMsg, err := srv.Do(&req)
	if err != nil {
		return nil, err
	}
	parts := CurlySplit(peekMsg)
	if parts[0] != "ok" {
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

// Poll performs an ACARS Poll (fetch with update)
//
// messages contains the list of responses provided by the server
func (srv *Server) Poll() (messages []*Message, err error) {
	req := Request{Logon: srv.Logon, From: srv.StationName, To: srv.StationName, Type: MsgTypePoll}
	pollMsg, err := srv.Do(&req)
	if err != nil {
		return nil, err
	}
	parts := CurlySplit(pollMsg)
	if parts[0] != "ok" {
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
