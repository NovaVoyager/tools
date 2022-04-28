package tools

import "time"

const (
	//TimeDateFormat 时间日期格式
	TimeDateFormat = "2006-01-02 15:04:05"

	//DateFormat 日期格式
	DateFormat = "2006-01-02"
)

var (
	//cstZone 东八区时区
	cstZone = time.FixedZone("CST", 3600*8)

	//TimeStr 时间格式化方法
	TimeStr = func(format string, cst *time.Location) string {
		return time.Now().In(cst).Format(format)
	}
)
