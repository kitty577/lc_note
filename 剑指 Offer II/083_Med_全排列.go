package main

func permute(nums []int) [][]int {
	res := [][]int{}
	arr := []int{}

	length := len(nums)
	used := make([]bool, length)

	var find func()
	find = func() {
		if len(arr) == length { // 集齐了
			res = append(res, append([]int{}, arr...))
			return
		}

		for i := 0; i < length; i++ {
			if used[i] {
				continue
			}

			// 选
			arr = append(arr, nums[i])
			used[i] = true
			find()

			// 还原现场
			arr = arr[:len(arr)-1]
			used[i] = false
		}
	}

	find()
	return res
}
