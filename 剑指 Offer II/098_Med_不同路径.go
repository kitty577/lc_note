package main

// 动态规划，dp[i][j]为走到[i][j]单元格的路径数
// 只能向下 和 向右走
// i==0 ==> dp[0][j]=1
// j==0 ==> dp[i][0]=1
// i>0&&j>0  ==> dp[i][j]=dp[i-1][j]+dp[i][j-1]

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
