package main

// 递归
// 对每一个元素而言 有选和不选两种
func subsets(nums []int) [][]int {
	ans := [][]int{} // 存放最终答案
	arr := []int{}   // 子集
	var find func(int)
	find = func(curIndex int) {
		// 超出边界
		if curIndex == len(nums) {
			ans = append(ans, append([]int{}, arr...))
			return
		}
		// 选该值
		arr = append(arr, nums[curIndex])
		find(curIndex + 1)
		// 还原现场
		arr = arr[:len(arr)-1]

		// 不选该值
		find(curIndex + 1)
	}

	find(0)
	return ans
}

// 2022.11.20
