package main

// 下标是跳跃的最大长度 也就是说其实是一个可跳跃的范围，不用纠结一次跳几步

//  问题也就转化为 跳跃的范围是否能覆盖到终点
// 因此可以采用贪心，每次取最大距离，达到结果最大
func canJump(nums []int) bool {
	cover := 0
	n := len(nums) - 1
	for i := 0; i <= cover; i++ {
		cover = max(i+nums[i], cover)
		if cover >= n {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
