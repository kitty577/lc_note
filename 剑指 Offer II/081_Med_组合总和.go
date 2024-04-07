package main

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	arr := []int{}

	var find func(curIndex, curSum int)
	find = func(curIndex, curSum int) {
		// 集齐了
		if curSum == target {
			ans = append(ans, append([]int{}, arr...))
			return
		}

		// 超出边界了
		if curIndex == len(candidates) {
			return
		}

		// 不选
		find(curIndex+1, curSum)

		// 选
		if curSum+candidates[curIndex] <= target {
			curSum += candidates[curIndex]
			arr = append(arr, candidates[curIndex])
			// 由于一个元素可以选多次，因此下一次遍历，还是要从此元素开始
			find(curIndex, curSum)

			// 还原现场
			arr = arr[:len(arr)-1]
		}
	}

	find(0, 0)
	return ans
}
