package main

// 广度优先搜索

// 以坏橘子为中心 上下左右四个方向搜索，四个方向搜索完 time+1
// 可以记录下怀橘子的位置，队列存储；每次就取队列中的元素搜索腐烂操作
// 同时也要记录下新鲜橘子的个数  如果坏橘子队列中的元素都遍历完 还有新鲜橘子，说明不可能全腐烂完

// 方向数组 上下左右
var (
	dirsX = []int{0, 0, -1, 1}
	dirsY = []int{1, -1, 0, 0}
)

func orangesRotting(grid [][]int) int {
	row := len(grid)
	column := len(grid[0])
	fresh := 0
	bad := [][]int{}

	minTime := 0

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if grid[i][j] == 1 {
				fresh++
			} else if grid[i][j] == 2 {
				bad = append(bad, append([]int{}, i, j))
			}
		}
	}

	for len(bad) > 0 {
		badCount := len(bad)
		rotten := false
		for i := 0; i < badCount; i++ {
			// 当前处理bad[0]因此需要将bad[0]从bad队列中删除，更新，避免重复处理该点
			curPos := bad[0]
			bad = bad[1:]
			for k := 0; k < 4; k++ {
				x := curPos[0] + dirsX[k]
				y := curPos[1] + dirsY[k]
				if x >= 0 && x < row && y >= 0 && y < column && grid[x][y] == 1 {
					grid[x][y] = 2
					bad = append(bad, append([]int{}, x, y))
					fresh--
					rotten = true
				}
			}
		}
		if rotten {
			minTime++
		}
	}

	if fresh == 0 {
		return minTime
	}
	return -1
}

// 2022.12.24
