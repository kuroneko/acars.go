package acars

type MsgType string

const (
	MsgTypeProgress = MsgType("progress")
	MsgTypeCPDLC    = MsgType("cpdlc")
	MsgTypeTelex    = MsgType("telex")
	MsgTypePing     = MsgType("ping")
	MsgTypePosReq   = MsgType("posreq")
	MsgTypePosition = MsgType("position")
	MsgTypeDataReq  = MsgType("datareq")
	MsgTypePoll     = MsgType("poll")
	MsgTypePeek     = MsgType("peek")
	MsgTypeAdsC     = MsgType("ads-c")
)

func (mt MsgType) String() string {
	return string(mt)
}
