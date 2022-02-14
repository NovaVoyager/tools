package tools

import (
	"math/rand"
	"time"
)

//RandKind 随机种类
type RandKind int

const (
	KcRandKindNum   RandKind = 0 // 纯数字
	KcRandKindLower RandKind = 1 // 小写字母
	KcRandKindUpper RandKind = 2 // 大写字母
	KcRandKindAll   RandKind = 3 // 数字、大小写字母
)

//Krand 随机字符串
func Krand(size int, kind RandKind) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = RandKind(rand.Intn(3))
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}
