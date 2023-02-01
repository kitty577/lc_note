package main

// 纯模拟螺旋顺序
// Attention: 边界！！！

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	res := []int{}
	row, column := len(matrix), len(matrix[0])
	left, right, top, bottom := 0, column-1, 0, row-1

	// 模拟螺旋顺序

	for left <= right && top <= bottom {
		// 从左到右(含最右)
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}

		// 从右到下(含最下)
		for j := top + 1; j <= bottom; j++ {
			res = append(res, matrix[j][right])
		}

		if left < right && top < bottom {
			// 从右到左(含最左)
			for m := right - 1; m >= left; m-- {
				res = append(res, matrix[bottom][m])
			}

			// 从左到上
			for n := bottom - 1; n > top; n-- {
				res = append(res, matrix[n][left])
			}
		}
		left++
		right--
		bottom--
		top++
	}
	return res
}

// 2022.12.16
