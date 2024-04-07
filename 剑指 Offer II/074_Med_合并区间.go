package main

import "sort"

// 左取小，右取大
func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := [][]int{intervals[0]}
	left, right := 0, 0

	for i := 1; i < len(intervals); i++ {
		pre := ans[len(ans)-1]
		if intervals[i][0] > pre[1] { // 独立两个区间，没有交集
			ans = append(ans, intervals[i])
		} else {
			ans = ans[:len(ans)-1]
			left = min(pre[0], intervals[i][0])
			right = max(pre[1], intervals[i][1])
			ans = append(ans, []int{left, right})
		}
	}
	return ans
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func min[T Ordered](a, b T) T {
	if a > b {
		return b
	}
	return a
}

func max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
