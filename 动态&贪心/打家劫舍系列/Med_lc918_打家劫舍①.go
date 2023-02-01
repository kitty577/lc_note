package main

// 偷到的最高金额---想到动态规划
// 相邻两间屋子不能偷！！

// dp[i]表示偷第i间屋子时 能达到的最高金额

// 首先考虑最简单的情况，只有一间屋子，那么偷它 即达到最高金额,dp[0]=nums[0]
// 如果有两间，由于相邻规则，只能偷其中一间,dp[1]=max(nums[0],nums[1])
// 如果大于两间，那么对于第i间（i>2)而言，有两个选择：
//	(1)偷：偷第i间，那么不能偷第i-1间，dp[i]=dp[i-2]+nums[i]
// 	(2)不偷：不偷第i间，那么偷第i-1间，dp[i]=dp[i-1]

// =====>> dp[i]=max(dp[i-2]+nums[i],dp[i-1]) (i>2)

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 很明显，需要记录的只有dp[i-2]和dp[i-1]这两个状态，那么可以使用变量记录，而不用开辟dp数组
func rob2(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}

	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < length; i++ {
		tmp := second
		second = max(first+nums[i], second)
		first = tmp
	}
	return second

}

// 2022.12.14
