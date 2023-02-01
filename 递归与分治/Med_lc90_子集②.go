package main

import "sort"

// nums中存在重复元素。但结果要求不能包含重复子集
// 与78.子集类似，对于nums中nums[i]而且，有 “选” or “不选” 两种可能
// 造成重复子集的原因:nums相同的元素的选和不选
// 如 [1,2,2]中，[1,2(第一个2被选，第二个2不选)]和[1,2(第一个2不选，第二个2选)]
func subsetsWithDup(nums []int) [][]int {
	// 先对nums排序 (使得相同的元素相邻)
	sort.Ints(nums)

	res := [][]int{}
	arr := []int{}
	var find func(bool, int)
	find = func(chose bool, curIndex int) {
		// 边界 退出递归
		if curIndex == len(nums) {
			res = append(res, append([]int{}, arr...))
			return
		}
		// 不选
		find(false, curIndex+1)

		// 选
		// 这里要注意去重！
		if curIndex > 0 && nums[curIndex-1] == nums[curIndex] && !chose {
			return
		}
		arr = append(arr, nums[curIndex])

		find(true, curIndex+1)

		// 还原现场
		arr = arr[:len(arr)-1]
	}
	find(false, 0)
	return res
}

// 2022.11.30
