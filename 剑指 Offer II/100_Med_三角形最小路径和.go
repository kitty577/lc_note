package main

// 动态规划 n:=len(triangle) dp[n][n]
// dp[i][j]表示[i][j]位置的路径和
// 2
// 3 4
// 6 5 7
// 4 1 8 3

// i==0&&j==0  --> dp[0][0]=triangle[0][0]
// 由于题中标明的走向，那么可知，一个位置的路径只能来源于正上方 和 左上方
//
//	i>0&&j==0  --> 即对于第一列来说，没有左上方，只有正上方 dp[i][0]=dp[i-1][0]
//	i>0&&j>0  --> dp[i][j]=min(dp[i-1][j],dp[i-1][j-1]) + triangle[i][j]
func minimumTotal(triangle [][]int) int {
	length := len(triangle)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return triangle[0][0]
	}

	dp := make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, length)
	}

	dp[0][0] = triangle[0][0]
	for i := 1; i < length; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}

	ans := dp[length-1][0]
	for j := 1; j < length; j++ {
		if ans > dp[length-1][j] {
			ans = dp[length-1][j]
		}
	}

	return ans
}
