package main

// 动态规划
// 创建dp[n]数组，其中dp[i]为偷至i家时累计的财富
// 状态转移方程为:dp[i]=max(dp[i-1],dp[i-2]+nums[i])   //  上一家偷了，那么这家不能偷； 上一家没偷，则偷这家
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	} else if n == 2 {
		return max(nums[0], nums[1])
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 在上面的基础上，其实状态转移只需要 上一个、上上个，两个变量维护就行，可以省去数组
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	} else if n == 2 {
		return max(nums[0], nums[1])
	}

	first := nums[0]
	second := max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		tmp := second
		second = max(first+nums[i], second)
		first = tmp
	}

	return second
}
