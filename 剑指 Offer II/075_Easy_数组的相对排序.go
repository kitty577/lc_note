package main

import "sort"

// 可以先统计arr1中数字数量，按arr2的顺序放入新数组
func relativeSortArray(arr1 []int, arr2 []int) []int {
	counts := make(map[int]int)
	for _, num := range arr1 {
		counts[num]++
	}

	ans := []int{}
	for _, n := range arr2 {
		tmp := []int{}
		for i := 0; i < counts[n]; i++ {
			tmp = append(tmp, n)
		}
		ans = append(ans, tmp...)
		delete(counts, n)
	}

	// 处理剩余的数
	last := []int{}
	for value, times := range counts {
		for i := 0; i < times; i++ {
			last = append(last, value)
		}
	}

	sort.Ints(last)

	ans = append(ans, last...)
	return ans
}
