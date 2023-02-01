package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	arr := []int{}
	var used = make([]bool, len(candidates))
	sum := 0
	var find func(int)
	find = func(index int) {
		// 满足条件 退出递归
		if sum == target {
			res = append(res, append([]int{}, arr...))
			return
		}

		// 剪枝
		if sum > target || index == len(candidates) || used[index] {
			return
		}

		// 不选
		find(index + 1)

		// 选
		// 去重
		if index > 0 && !used[index-1] && candidates[index] == candidates[index-1] {
			return
		}

		arr = append(arr, candidates[index])
		sum = sum + candidates[index]
		used[index] = true
		find(index + 1)

		// 还原现场
		arr = arr[:len(arr)-1]
		sum = sum - candidates[index]
		used[index] = false
	}
	find(0)
	return res

}

// 2022.12.07
