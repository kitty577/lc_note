// 动态规划
// 创建n+1长度的数组dp,其中dp[i]表示到下标i的最小花费
// 当 2 <= i <=n时，可以从下标 i-1使用 cost[i-1]的花费达到下标i，
// 或者从下标 i-2使用cost[i−2]的花费达到下标 i。

// 为了使总花费最小，dp[i] 应取上述两项的最小值，因此状态转移方程如下：
// dp[i]=min(dp[i-1]+cost[i-1],dp[i-2]+cost[i-2])

package main

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
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

// 2022.10.09
