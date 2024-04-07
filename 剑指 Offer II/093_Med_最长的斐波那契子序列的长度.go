package main

// 斐波那契 需要记录至少需要记录两个数才可以推断序列
// 定义dp数组，dp[i][j]表示以arr[i]、arr[j]作为最后两个数字的斐波那契数列的长度

func lenLongestFibSubseq(arr []int) int {
	ans := 0
	lenghth := len(arr)

	dp := make([][]int, lenghth)
	for i := 0; i < lenghth; i++ {
		dp[i] = make([]int, lenghth)
	}

	indexMap := make(map[int]int, lenghth) // 元素值---下标
	for idx, x := range arr {
		indexMap[x] = idx
	}

	for i := 0; i < lenghth; i++ {
		for j := i + 1; j < lenghth; j++ {
			dp[i][j] = max(dp[i][j], 2)
			nextValue := arr[i] + arr[j]
			if nextIdx, ok := indexMap[nextValue]; ok {
				dp[j][nextIdx] = dp[i][j] + 1
				ans = max(ans, dp[j][nextIdx])
			}

		}
	}
	return ans
}
