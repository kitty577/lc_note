package main

// 回溯
// 每个元素都有两种选择 被加 / 被减
func findTargetSumWays(nums []int, target int) int {
	var res int
	var backtracking func(curIdx int, sum int)
	backtracking = func(curIdx int, sum int) {
		if curIdx == len(nums) { // 已经处理完了数组中的全部元素
			if sum == target {
				res++
			}
			return
		}
		backtracking(curIdx+1, sum+nums[curIdx])
		backtracking(curIdx+1, sum-nums[curIdx])
	}

	backtracking(0, 0)

	return res
}
