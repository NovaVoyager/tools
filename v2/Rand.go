package tools

import (
	"math/rand"
	"time"
)

//GetRandRang 获取指定范围内的随机数
func GetRandRang(min, max int) int {
	if min > max {
		return 0
	}
	rand.Seed(time.Now().UnixNano())
	res := rand.Intn(max-min) + min
	return res
}
