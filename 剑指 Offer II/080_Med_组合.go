package main

// 对于每一个数都有 选 和 不选 两种
func combine(n int, k int) [][]int {
	ans := [][]int{}
	arr := []int{}

	var find func(cur int)
	find = func(cur int) {
		if len(arr) == k { // 已经收集满了
			ans = append(ans, append([]int{}, arr...))
			return
		}

		// 就算全选上 也不够了; 或者 已经超出可选范围了
		if len(arr)+n-cur+1 < k || cur == n+1 {
			return
		}

		// 不选
		find(cur + 1)

		// 选
		arr = append(arr, cur)
		find(cur + 1)

		// 还原现场
		arr = arr[:len(arr)-1]
	}

	find(1)

	return ans
}
