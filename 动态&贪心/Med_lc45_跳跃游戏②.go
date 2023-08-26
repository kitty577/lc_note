package main

// 跳跃1 求是否可达，只要考虑范围是否能覆盖终点
// 跳跃2 求到达终点的最小跳跃次数 ===> 跳跃的长度要够大 则次数少
// 还是可以从覆盖范围出发，如果当前的最远距离覆盖不了终点，那就步数+1
func jump(nums []int) int {
	length := len(nums)
	if length == 1 {
		return 0
	}

	end := 0         // 当前位置的可覆盖范围的右边界
	maxPosition := 0 // 最远覆盖范围
	step := 0

	for i := 0; i < length-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		for i == end {
			// 到达原边界了 更新边界
			end = maxPosition
			step++
		}
	}

	return step
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
