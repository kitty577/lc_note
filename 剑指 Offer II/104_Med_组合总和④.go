package main

// 动态规划 dp[target+1]
// dp[i]表示选取的元素之和为i的方案数
// i=0时，也即和为0,那么只有一种方案：不选任何元素
// 当i>0时， 该排列的最后一个元素一定来自于数组，假设该元素为num，那么此时的排列数其实就转化为 和为i-num的排列后面加一个num元素。
// 因此计算dp[i]要计算 dp[i-num]的值
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for j := 1; j <= target; j++ {
		for _, num := range nums {
			if num <= j {
				dp[j] += dp[j-num]
			}
		}
	}
	return dp[target]
}
