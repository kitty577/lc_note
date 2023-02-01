// 动态规划
// 机器人每次只能向下或者向右移动一步====》值只能来自于上方和左方
// 设dp[i][j]为走到单元格[i][j]累计的不同路径数
// 当i==0&&j==0 f[i][j]=0
// 当i==0&&j>0 f[i][j]=f[i][j-1]  第一行元素只能从左边获得
// 当i>0&&j==0 f[i][j]=f[i-1][j]  第一列元素只能从上边获得
// 当i>0&&j>0  f[i][j]=f[i-1][j]+f[i][j-1]

package main

func uniquePaths(m int, n int) int {
	var dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
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

//2022.10.09
