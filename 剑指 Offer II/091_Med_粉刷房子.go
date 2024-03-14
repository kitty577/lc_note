package main

// 动态规划
// dp[i][j]为使用j颜料粉刷i号房子时的最小花费
// 状态转移方程：dp[i][j]=min(dp[i-1][j+1]mod3,dp[i-1][j+2]mod3)+cos[i][j]

func minCost(costs [][]int) int {
	row := len(costs)
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, 3)
		if i == 0 {
			dp[i] = costs[i]
		}
	}

	for i := 1; i < row; i++ {
		for j := 0; j < 3; j++ {
			dp[i][j] = min(dp[i-1][(j+1)%3], dp[i-1][(j+2)%3]) + costs[i][j]
		}
	}

	return min(min(dp[row-1][0], dp[row-1][1]), dp[row-1][2])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
