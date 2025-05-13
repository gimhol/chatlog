package opts

import "time"

type OptsGetMessages struct {
	Asc       bool
	StartTime time.Time
	EndTime   time.Time
	Talker    string
	Sender    string
	Keyword   string
	Limit     int
	Offset    int
}

func NewOptsGetMessages() *OptsGetMessages {
	ret := &OptsGetMessages{
		Asc: true,
	}
	return ret
}

func (o *OptsGetMessages) Clone() *OptsGetMessages {
	ret := &OptsGetMessages{
		Asc:       o.Asc,
		StartTime: o.StartTime,
		EndTime:   o.StartTime,
		Talker:    o.Talker,
		Sender:    o.Sender,
		Keyword:   o.Keyword,
		Limit:     o.Limit,
		Offset:    o.Offset,
	}
	return ret
}
