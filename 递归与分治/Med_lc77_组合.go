package main

// 递归
func combine(n int, k int) [][]int {
	res := [][]int{}
	arr := []int{}
	var find func(int)
	find = func(index int) {
		// 已经选了超过K个或者选不满K个
		if len(arr) > k || len(arr)+n-index+1 < k {
			return
		}
		// 选到数组末尾了
		if index == n+1 {
			res = append(res, append([]int{}, arr...))
			return
		}

		// 不选
		find(index + 1)

		// 选
		arr = append(arr, index)

		find(index + 1)

		// 还原现场
		arr = arr[:len(arr)-1]
	}
	find(1)
	return res
}

// 2022.12.4
