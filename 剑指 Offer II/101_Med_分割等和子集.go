package main

// 数组分割成两部分，元素和相等
// 背包问题：0-1背包（背包容量为sum/2,物品为数组中的元素）
func canPartition(nums []int) bool {
	length := len(nums)
	if length < 2 {
		return false
	}

	sum, maxNum := 0, nums[0]
	for _, v := range nums {
		sum += v
		if v > maxNum {
			maxNum = v
		}
	}

	if sum%2 != 0 { // 和是奇数，无法平分
		return false
	}
	target := sum / 2
	if target < maxNum {
		return false
	}

	// dp[i][j]表示数组的[0...i)区间内的元素能否凑成和为j
	// 当i=0,也就是只有数组的第一个元素可供选择，因此能凑成和为 0 或 nums[0]
	// 当j=0,也就是所有一个元素都不选

	dp := make([][]bool, length)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	for i := 0; i < length; i++ {
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	for i := 1; i < length; i++ {
		v := nums[i]
		for j := 1; j <= target; j++ {
			if v > j { // 这个元素不能选
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i][j] || dp[i-1][j-v]
			}

		}
	}
	return dp[length-1][target]
}
