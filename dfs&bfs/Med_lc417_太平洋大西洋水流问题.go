package main

// 反向搜索

// 雨水的流动方向是从高到低，每个单元格的水只能流到高度小于或等于当前单元格的位置
// 从一个单元格开始 通过搜索的方式可以模拟水的流动

// 如果直接以每个单元格作为起点模拟雨水的流动，则会重复遍历每个单元格，导致时间复杂度过高。
// 为了降低时间复杂度，可以从矩阵的边界开始反向搜索寻找雨水流向边界的单元格，反向搜索时，每次只能移动到高度相同或更大的单元格

// 方位
var dirs = []struct {
	x, y int
}{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func pacificAtlantic(heights [][]int) (ans [][]int) {
	row, colum := len(heights), len(heights[0])
	pacific := make([][]bool, row)  // 记录网格内单元格能否流入太平洋
	atlantic := make([][]bool, row) // 记录网格内单元格能否流入大西洋
	for i := range pacific {
		pacific[i] = make([]bool, colum)
		atlantic[i] = make([]bool, colum)
	}

	var dfs func(int, int, [][]bool)
	dfs = func(x, y int, ocean [][]bool) {
		if ocean[x][y] {
			return
		}
		ocean[x][y] = true
		for _, d := range dirs {
			if nx, ny := x+d.x, y+d.y; nx >= 0 && nx < row && ny >= 0 && ny < colum && heights[nx][ny] >= heights[x][y] {
				dfs(nx, ny, ocean)
			}
		}
	}

	// 左边界
	for i := 0; i < row; i++ {
		dfs(i, 0, pacific)
	}

	// 上边界
	for j := 0; j < colum; j++ {
		dfs(0, j, pacific)
	}

	// 右边界
	for i := 0; i < row; i++ {
		dfs(i, colum-1, atlantic)
	}

	// 下边界
	for j := 0; j < colum; j++ {
		dfs(row-1, j, atlantic)
	}

	for i, s := range pacific {
		for j, flag := range s {
			if flag && atlantic[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return
}
