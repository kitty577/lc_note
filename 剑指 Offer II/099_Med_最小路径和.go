package main

// 动态规划 dp[i][j]表示走到[i][j]单元格的最小路径和
// i==0 ==> dp[0][j]=dp[0][j-1]+grid[0][j]
// j==0 ==> dp[i][0]=dp[i-1][0]+grid[i][0]
// i>0&&j>0 ==> dp[i][j]=min(dp[i-1][j],dp[i][j-1])+grid[i][j]
func minPathSum(grid [][]int) int {
	row := len(grid)
	column := len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if i == 0 && j > 0 {
				grid[i][j] += grid[i][j-1]
			}
			if i > 0 && j == 0 {
				grid[i][j] += grid[i-1][j]
			}
			if i > 0 && j > 0 {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[row-1][column-1]
}
