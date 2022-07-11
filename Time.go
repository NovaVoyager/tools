package tools

import (
	"time"
)

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

//SecondToDayHMS 秒转换为天 时 分 秒
func SecondToDayHMS(second int64) (int64, int64, int64, int64) {
	if second == 0 {
		return 0, 0, 0, 0
	}

	oneHour := 60 * 60
	oneMinute := 60
	oneDay := oneHour * 24

	// 天    总秒数/一天的秒数=几天
	d := second / int64(oneDay)
	// 小时  不够一天的秒数可以换算成几小时
	h := second / int64(oneHour)
	//h1 := second % int64(oneDay) / int64(oneHour)
	// 分钟  不够一小时的秒数可以换算成几分钟
	m := second % int64(oneHour) / int64(oneMinute)
	//秒数	不够一分钟的秒数可以换算成几秒
	s := second % int64(oneHour) % int64(oneMinute) % int64(oneMinute)

	return d, h, m, s
}
