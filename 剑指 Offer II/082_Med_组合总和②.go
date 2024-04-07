package main

// 提供的数组中含有重复元素，且结果集不能包含重复组合
// 想到map检验唯一存在， / 也可先将数组排序，在选择时就拦截重复
func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	ans := [][]int{}
	arr := []int{}
	used := make([]bool, len(candidates))
	var find func(curIndex, curSum int)
	find = func(curIndex, curSum int) {
		if curSum == target { // 集齐了
			ans = append(ans, append([]int{}, arr...))
			return
		}

		// 超出边界了、元素已经被使用了、和超出目标值了
		if curIndex == len(candidates) || used[curIndex] || curSum > target {
			return
		}

		// 不选
		find(curIndex+1, curSum)

		// 选

		if curIndex > 0 && candidates[curIndex] == candidates[curIndex-1] && !used[curIndex-1] {
			return
		}

		if curSum+candidates[curIndex] <= target {
			arr = append(arr, candidates[curIndex])
			used[curIndex] = true
			find(curIndex+1, curSum+candidates[curIndex])

			// 还原现场
			arr = arr[:len(arr)-1]
			used[curIndex] = false
		}
	}

	find(0, 0)
	return ans
}
