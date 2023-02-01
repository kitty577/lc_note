// 球的每一次运动都是从上一行到下一行的转移
// 1.卡在当前行 2.下一行左 3.下一行右

package main

// 球是从上一行到下一行的之间的转移
// 对球而言 有下左、下右、卡在当前行 这么几种状态
func findBall(grid [][]int) []int {
	row := len(grid)
	column := len(grid[0])
	ans := make([]int, column)

	// 处理每一列位置上的球
	for k := 0; k < column; k++ {
		// 初始值为-1（也即不可达）
		ans[k] = -1
		// curL为当前列
		curL := k

		// 球的行转移
		for i := 0; i < row; i++ {
			if grid[i][curL] == 1 {
				if curL < column-1 && grid[i][curL+1] == 1 { // 向右
					curL++
				} else {
					break
				}
			} else {
				if curL-1 >= 0 && grid[i][curL-1] == -1 { // 向左
					curL--
				} else {
					break
				}
			}
			if i == row-1 {
				ans[k] = curL
			}
		}
	}
	return ans
}

// 2022.10.19
