package main

func permute(nums []int) [][]int {
	res := [][]int{}                // 排列集合，最终的返回结果
	arr := []int{}                  // 一次全排列的结果
	used := make([]bool, len(nums)) // 记录nums[i]元素是否被使用
	var find func(nums []int, arr []int, used []bool)
	find = func(nums []int, arr []int, used []bool) {
		length := len(nums)
		if len(arr) == length {
			res = append(res, append([]int{}, arr...))
			return
		}

		for i := 0; i < length; i++ {
			if used[i] { // 该数已经选过了
				continue
			}

			// 选
			arr = append(arr, nums[i])
			used[i] = true
			find(nums, arr, used)

			// 还原现场
			arr = arr[:len(arr)-1]
			used[i] = false
		}
	}
	find(nums, arr, used)
	return res
}

// 函数参数：若为引用类型（map\channel\slice）可不用&取地址，值类型则需要传&类型，否则改动不了

// 2020.12.05
