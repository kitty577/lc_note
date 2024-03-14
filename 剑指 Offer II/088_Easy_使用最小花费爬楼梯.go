package main

// 动态规划
// 创建n+1长度的数组dp,dp[i]表示到达该阶梯需要耗费的最小体力
// 状态转移方程：dp[i]=min(dp[i-1]+cost[i-1],dp[i-2]+cost[i-2])

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	// 由于在开始时，可以选择从下标0 或 1 作为初始阶段，则到达0阶梯 和 到达1阶梯耗费的体力都为0
	dp[0] = 0
	dp[1] = 0
	for i := 2; i < n+1; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
