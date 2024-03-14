package main

// 房屋围城环， 有首无尾 、 无首有尾
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	} else if n == 2 {
		return max(nums[0], nums[1])
	}

	return max(listRob(nums[:n-1]), listRob(nums[1:]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func listRob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	} else if n == 2 {
		return max(nums[0], nums[1])
	}
	first := nums[0]
	second := max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		tmp := second
		second = max(first+nums[i], second)
		first = tmp
	}
	return second
}
