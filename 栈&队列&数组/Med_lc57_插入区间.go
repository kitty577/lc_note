package main

func insert(intervals [][]int, newInterval []int) [][]int {
	left, right := newInterval[0], newInterval[1]
	merged := false

	res := [][]int{}

	for _, interval := range intervals {
		if interval[0] > right { // 在插入区间的右侧 无交集
			if !merged {
				res = append(res, []int{left, right})
				merged = true // 处理完
			}
			res = append(res, interval)
		} else if interval[1] < left { // 在插入区间的左侧 无交集
			res = append(res, interval)
		} else { // 有交集
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}

	if !merged {
		res = append(res, []int{left, right})
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 2023.01.03
