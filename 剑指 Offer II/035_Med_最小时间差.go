package main

import (
	"math"
	"sort"
)

// 可以将时间转化为分钟表示s
func findMinDifference(timePoints []string) int {
	sort.Strings(timePoints)
	ans := math.MaxInt32
	t0Minute := getMinutes(timePoints[0])
	pre := t0Minute
	for _, t := range timePoints[1:] {
		tm := getMinutes(t)
		ans = min(ans, tm-pre)
		pre = tm
	}

	// 还要比较下首尾两个时间，跨零点情况
	ans = min(ans, t0Minute+1440-pre)
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getMinutes(t string) int {
	return (int(t[0]-'0')*10+int(t[1]-'0'))*60 + int(t[3]-'0')*10 + int(t[4]-'0')
}
