package tools

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestIntersection(t *testing.T) {
	type args[T interface{ int | int64 | uint64 }] struct {
		a []T
		b []T
	}
	type testCase[T interface{ int | int64 | uint64 }] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "返回两个切片的交集1",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				b: []int{3, 5, 9},
			},
			want: []int{3, 5, 9},
		},
		{
			name: "返回两个切片的交集2",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				b: []int{99, 77, 88, 66},
			},
			want: []int{},
		},
		{
			name: "返回两个切片的交集3",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				b: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersection(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifference(t *testing.T) {
	type args[T interface{ int | uint64 | int64 | string }] struct {
		a []T
		b []T
	}
	type testCase[T interface{ int | uint64 | int64 | string }] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "数组差集1",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				b: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want: []int{0},
		},
		{
			name: "数组差集2",
			args: args[int]{
				a: []int{99, 85, 12, 30, 45},
				b: []int{1, 5, 6, 7, 8, 9},
			},
			want: []int{99, 85, 12, 30, 45},
		},
		{
			name: "数组差集3",
			args: args[int]{
				a: []int{99, 85, 12, 30, 45},
				b: []int{99, 85, 12, 30, 45},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Difference(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsElemInSliceHead(t *testing.T) {
	type args[T interface{ int | int64 | uint64 }] struct {
		dst []T
		src []T
	}
	type testCase[T interface{ int | int64 | uint64 }] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "切片头部插入元素1",
			args: args[int]{
				dst: []int{1, 2, 3},
				src: []int{0},
			},
			want: []int{0, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsElemInSliceHead(tt.args.dst, tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsElemInSliceHead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsExistArr(t *testing.T) {
	type args[T interface{ int | int64 | uint64 }] struct {
		val T
		arr []T
	}
	type testCase[T interface{ int | int64 | uint64 }] struct {
		name  string
		args  args[T]
		want  bool
		want1 int
	}
	tests := []testCase[int]{
		{
			name: "判断一个值是否在数组中存在1",
			args: args[int]{
				val: 0,
				arr: []int{0, 1, 2, 3},
			},
			want:  true,
			want1: 0,
		},
		{
			name: "判断一个值是否在数组中存在2",
			args: args[int]{
				val: 99,
				arr: []int{0, 1, 2, 3},
			},
			want:  false,
			want1: -1,
		},
		{
			name: "判断一个值是否在数组中存在3",
			args: args[int]{
				val: 0,
				arr: []int{},
			},
			want:  false,
			want1: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := IsExistArr(tt.args.val, tt.args.arr)
			if got != tt.want {
				t.Errorf("IsExistArr() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("IsExistArr() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRemoveRepeat(t *testing.T) {
	type args[T interface{ int | uint64 | string }] struct {
		arr []T
	}
	type testCase[T interface{ int | uint64 | string }] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "数组去重1",
			args: args[int]{
				arr: []int{1, 1, 2, 2, 3, 5, 0, 0, 5, 6, 7, 7, 3},
			},
			want: []int{1, 2, 3, 5, 0, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveRepeat(tt.args.arr)
			sort.Slice(got, func(i, j int) bool {
				return got[i] < got[j]
			})

			sort.Slice(tt.want, func(i, j int) bool {
				return tt.want[i] < tt.want[j]
			})

			fmt.Println(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveRepeat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveRepeatInSitu(t *testing.T) {
	type args[T interface{ int | uint64 }] struct {
		arr []T
	}
	type testCase[T interface{ int | uint64 }] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "原地数组去重1",
			args: args[int]{
				arr: []int{1, 2, 2, 3, 4, 5, 4, 7},
			},
			want: []int{1, 2, 3, 4, 5, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveRepeatInSitu(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveRepeatInSitu() = %v, want %v", got, tt.want)
			}
		})
	}
}
