package main

// 由于是升序排列 然后再被翻转
// 二分查找
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[0] {
			if nums[mid] > target && nums[0] <= target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && nums[len(nums)-1] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
