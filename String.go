package tools

import (
	"math/rand"
	"strings"
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

//KRand 随机字符串
func KRand(size int, kind RandKind) string {
	iKind, kinds, result := kind, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			iKind = RandKind(rand.Intn(3))
		}
		scope, base := kinds[iKind][0], kinds[iKind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

// FirstUpper 字符串首字母大写
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

// StrEscape 字符串转义
// 转义字符串里的 单引号 和 双引号
func StrEscape(str string) string {
	if str == "" {
		return ""
	}

	s := strings.ReplaceAll(str, "'", `\'`)
	s = strings.ReplaceAll(s, `"`, `\"`)

	return s
}
