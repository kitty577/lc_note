package main

// 对于每一个元素 都有 选和不选 两种可能
func subsets(nums []int) [][]int {
	ans := [][]int{}
	arr := []int{}
	var find func(curIndex int)
	find = func(curIndex int) {
		if curIndex == len(nums) {
			ans = append(ans, append([]int{}, arr...))
			return
		}

		// 不选
		find(curIndex + 1)

		// 选
		arr = append(arr, nums[curIndex])
		find(curIndex + 1)

		// 还原现场
		arr = arr[:len(arr)-1]
	}

	find(0)
	return ans
}
