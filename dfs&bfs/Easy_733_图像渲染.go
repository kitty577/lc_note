// dfs
// 以image[sr][sc]为起始点 上下左右四个方位dfs 注意要存储原始位置的像素值
// 方位数组

package main

var (
	dx = []int{-1, 1, 0, 0}
	dy = []int{0, 0, -1, 1}
)

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	currColor := image[sr][sc]
	if currColor != color {
		dfs(image, sr, sc, currColor, color)
	}
	return image
}

func dfs(image [][]int, x, y, currColor, color int) {
	if image[x][y] == currColor {
		image[x][y] = color // 如果像素相同，需要上色填充
		for i := 0; i < 4; i++ {
			newX, newY := x+dx[i], y+dy[i]
			if newX >= 0 && newX < len(image) && newY >= 0 && newY < len(image[0]) { // 越界限
				dfs(image, newX, newY, currColor, color)
			}
		}
	}
}

// 2022.10.10
