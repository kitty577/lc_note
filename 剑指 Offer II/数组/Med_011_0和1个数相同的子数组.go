package main

// 和上一题的解法类似
// 0,1个数相同 ==》这里需要把0转化为-1
// 转化后，如果一段区间内0和1数量相同   ===》 这段区间内的和为0
func findMaxLength(nums []int) int {
	pre, ans := 0, 0
	m := make(map[int]int, 0) // key 为前缀和，value为下标
	m[0] = -1
	for i, v := range nums {
		if v == 1 {
			pre++
		} else {
			pre--
		}

		if index, ok := m[pre]; ok {
			ans = max(ans, i-index)
		} else {
			m[pre] = i
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
