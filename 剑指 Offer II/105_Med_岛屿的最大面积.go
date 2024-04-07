package main

// 深度优先
// 对于[i,j]==1时，深度优先搜索[i+1,j],[i-1,j],[i,j+1],[i,j-1]
// 剪枝： 走到边界; 值为0
func maxAreaOfIsland(grid [][]int) int {
	var res int
	var dfs func(i, j int, area *int)
	dfs = func(i, j int, area *int) {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == 0 {
			return
		}

		*area++

		// 已经累计了这个位置，置为0 防止重复累加
		grid[i][j] = 0

		// 上下左右四个方向搜索
		dfs(i-1, j, area)
		dfs(i+1, j, area)
		dfs(i, j-1, area)
		dfs(i, j+1, area)
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				temp := 0
				dfs(i, j, &temp)
				if temp > res {
					res = temp
				}
			}
		}
	}
	return res
}
