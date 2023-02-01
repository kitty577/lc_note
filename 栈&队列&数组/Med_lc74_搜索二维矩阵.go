package main

// 每行是从左到右升序，后面的行里的数大于前一行
func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	column := len(matrix[0])
	// 小于最小值 或者大于最大值
	if target < matrix[0][0] || target > matrix[row-1][column-1] {
		return false
	}

	var i int
	// 定位target处在哪一行
	for i = 0; i < row; {
		if matrix[i][column-1] < target { // target大于当前行的最大值，即继续找下一行
			i++
		} else {
			break
		}
	}

	// 找到了target的所在行数为i,开始对这一行二分查找
	left, right := 0, column-1
	for left <= right {
		mid := (left + right) / 2
		if matrix[i][mid] == target {
			return true
		} else if matrix[i][mid] > target {
			right = mid - 1
		} else if matrix[i][mid] < target {
			left = mid + 1
		}
	}
	return false
}

// 2022.12.22
