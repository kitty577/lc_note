package main

import "sort"

// 想到90.子集② 的做法，先排序，通过限制相邻元素是否选中来去重
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	res := [][]int{}
	arr := []int{}
	var used = make([]bool, length)
	var find func()
	find = func() {
		if len(arr) == len(nums) {
			res = append(res, append([]int{}, arr...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] || i > 0 && !used[i-1] && nums[i] == nums[i-1] {
				continue
			}

			// 选
			arr = append(arr, nums[i])
			used[i] = true

			// 下一次决策
			find()

			// 还原现场
			arr = arr[:len(arr)-1]
			used[i] = false
		}
	}
	find()
	return res
}

// 2022.12.05
