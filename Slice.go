package tools

import (
	"fmt"
	"math/rand"
	"time"
)

// RemoveRepeatInSitu 原地数组去重
func RemoveRepeatInSitu[T int | uint64](arr []T) []T {
	set := make(map[T]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}

	return arr[:j]
}

// RemoveRepeat 数组去重
func RemoveRepeat[T int | uint64](arr []T) []T {
	set := make(map[T]struct{}, len(arr))

	for _, v := range arr {
		set[v] = struct{}{}
	}

	newArr := make([]T, 0, len(set))

	for k, _ := range set {
		newArr = append(newArr, k)
	}

	return newArr
}

// ShuffleSlice 随机打乱切片
func ShuffleSlice[T uint64 | string](slice []T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	r.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

// IntToStr 整型切片转为字符串切片
func IntToStr[T uint64 | int64 | int](arr []T) []string {
	strArr := make([]string, 0, len(arr))
	for _, v := range arr {
		strArr = append(strArr, fmt.Sprintf("%d", v))
	}

	return strArr
}

// IsExistArr 判断一个值是否在数组中存在
// 存在返回 true 和 索引下标,不存在返回 false 和 -1 索引
func IsExistArr[T int | int64 | uint64](val T, arr []T) (bool, int) {
	tmpM := make(map[T]int, len(arr))

	for i, v := range arr {
		tmpM[v] = i
	}

	if index, ok := tmpM[val]; ok {
		return true, index
	}

	return false, -1
}

// InsElemInSliceHead 在切片头部插入元素
func InsElemInSliceHead[T int | int64 | uint64](dst, src []T) []T {
	newArr := make([]T, len(dst)+len(src))

	copy(newArr, src)
	copy(newArr[len(src):], dst)

	return newArr
}

// Intersection 返回两个切片的交集
func Intersection[T int | int64 | uint64](a, b []T) []T {
	result := make([]T, 0)
	m := make(map[T]bool)

	// 将切片 a 中的元素存入 map
	for _, value := range a {
		m[value] = true
	}

	// 遍历切片 b，如果元素在 map 中存在，则加入交集结果
	for _, value := range b {
		if m[value] {
			result = append(result, value)
		}
	}

	return result
}

// Difference 返回两个切片的差集（a 中有而 b 中没有的元素）
func Difference[T int | uint64 | int64 | string](a, b []T) []T {
	result := make([]T, 0)
	m := make(map[T]bool)

	// 将切片 b 中的元素存入 map
	for _, value := range b {
		m[value] = true
	}

	// 遍历切片 a，如果元素在 map 中不存在，则加入差集结果
	for _, value := range a {
		if !m[value] {
			result = append(result, value)
		}
	}

	return result
}
