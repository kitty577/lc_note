package main

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	arr := []int{}
	var find func(int, int)
	find = func(index int, ret int) {
		if len(arr) == k && ret == 0 {
			res = append(res, append([]int{}, arr...))
			return
		}
		// 如果剩下的数全选上也不够k个 或者 和已经超过n
		if len(arr)+9-index+1 < k || ret < 0 {
			return
		}

		// 不选
		find(index+1, ret)

		// 选
		arr = append(arr, index)
		find(index+1, ret-index)

		// 还原现场
		arr = arr[:len(arr)-1]
	}
	find(1, n)
	return res
}

// 2022.12.6
