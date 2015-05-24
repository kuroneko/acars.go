package acars

const (
	MsgTypeProgress = iota
  	MsgTypeCPDLC
  	MsgTypeTelex
  	MsgTypePing
  	MsgTypePosReq
  	MsgTypePosition
  	MsgTypeDataReq
  	MsgTypePoll
  	MsgTypePeek
  	MsgTypeAdsC
)

var msgTypeMap = []string{
	"progress",
	"cpdlc",
	"telex",
	"ping",
	"posreq",
	"position",
	"datareq",
	"poll",
	"peek",
	"ads-c",
}

func reqTypeToString(reqType int) string {
	if reqType < len(msgTypeMap) {
		return msgTypeMap[reqType]
	} else {
		return ""
	}
}

func reqTypeFromString(reqType string) int {
	for i, v := range msgTypeMap {
		if v == reqType {
			return i;
		}
	}
	return -1;
}
