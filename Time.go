package tools

import "time"

const (
	TimeNormalFormat = "2006-01-02 15:04:05"
)

var (
	cstZone = time.FixedZone("CST", 3600*8)
	TimeStr = func() string {
		return time.Now().In(cstZone).Format(TimeNormalFormat)
	}
)
