package main

// 二分
func peakIndexInMountainArray(arr []int) int {
	left, right := 1, len(arr)-2
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		} else if arr[mid] < arr[mid-1] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return 0
}
