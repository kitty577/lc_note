// 动态规划
// 每次只能向下或者向右移动一步====》值只能来自于上方和左方
// 设dp[i][j]为走到单元格[i][j]累计的最大价值
// 当i==0&&j==0 f[i][j]=0
// 当i==0&&j>0 f[i][j]=f[i][j-1]  第一行元素只能从左边获得
// 当i>0&&j==0 f[i][j]=f[i-1][j]  第一列元素只能从上边获得
// 当i>0&&j>0  f[i][j]=f[i-1][j]+f[i][j-1]

package main

func maxValue(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 && j > 0 { //第一行
				grid[i][j] += grid[i][j-1]
			} else if i > 0 && j == 0 { // 第一列
				grid[i][j] += grid[i-1][j]
			} else if i > 0 && j > 0 {
				grid[i][j] += max(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[row-1][col-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 2022.10.10
