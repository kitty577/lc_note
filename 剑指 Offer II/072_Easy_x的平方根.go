package main

// 二分查找
func mySqrt(x int) int {
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)/2
		tmp := mid * mid
		if tmp == x {
			return mid
		} else if tmp > x {
			right = mid - 1
		} else if tmp < x {
			left = mid + 1
		}
	}
	return right
}
