package main

// 二分查找
// 因为是有序，重复的数会在一起，那么在出现单个数之前，两个重复数的下标应该为 偶+奇 的组合，在出现单个数之后，下标变为 奇+偶的组合
// 可以根据下标的规律，划定单个数出现的范围
func singleNonDuplicate(nums []int) int {
	length := len(nums)
	if length == 0 {
		return nums[0]
	}
	left, right := 0, length-1
	for left <= right {
		mid := left + (right-left)/2
		if mid == 0 {
			left = mid + 1
			continue
		}
		if mid == length-1 {
			right = mid - 1
			continue
		}

		if mid%2 == 0 {
			if nums[mid] == nums[mid+1] { // 偶+奇
				left = mid + 1
			} else if nums[mid] == nums[mid-1] { // 奇+偶
				right = mid - 1
			} else {
				return nums[mid]
			}
		} else if mid%2 == 1 {
			if nums[mid] == nums[mid+1] { // 奇+偶
				right = mid - 1
			} else if nums[mid] == nums[mid-1] { // 偶+奇
				left = mid + 1
			} else {
				return nums[mid]
			}
		}
	}

	if right == 0 {
		return nums[right]
	}

	if left == length-1 {
		return nums[left]
	}

	return 0

}
