package main

// 对原数组排序
func permuteUnique(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	res := [][]int{}
	arr := []int{}

	lenght := len(nums)
	used := make([]bool, lenght)

	var find func()
	find = func() {
		if len(arr) == lenght {
			res = append(res, append([]int{}, arr...))
			return
		}

		for i := 0; i < lenght; i++ {
			if used[i] {
				continue
			}

			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			// 选
			used[i] = true
			arr = append(arr, nums[i])
			find()

			// 还原
			used[i] = false
			arr = arr[:len(arr)-1]
		}
	}

	find()
	return res
}
