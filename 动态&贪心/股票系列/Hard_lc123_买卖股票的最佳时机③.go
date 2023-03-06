package main

// 几个状态：
// 从未买入
// 第一次买入
// 第一次卖出
// 第二次买入
// 第二次卖出

func maxProfit(prices []int) int {
	// buy1,sell1,buy2,sell2 分别代表截止到当前状态的最大利润
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0

	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, prices[i]+buy1)
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}

	return sell2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp[i][j]:第i天，j=0~5个状态所获得最大利润
// j = 0: 不进行操作; j = 1: 第一次买入; j = 2：第一次卖出; j = 3: 第二次买入; j = 4: 第二次卖出

// dp[i][0] = dp[i-1][0]

// dp[i][1] = max(dp[i - 1][1], dp[i - 1][0] - prices[i])
// max(第i天没有操作沿用i-1天的状态（已经买过了，但可能选择下一次进行第一次买入花费最小），第i天第一次买入)

// dp[i][2] = max(dp[i - 1][2], dp[i - 1][1] + prices[i])
// max(第i天没有操作沿用i-1天的状态（已经卖出了，但可能选择下一次进行第一次卖出钱更多），第i天第一次卖出)

// dp[i][3] = max(dp[i - 1][3], dp[i - 1][2] - prices[i])
// max(第i天没有操作沿用i-1天的状态（已经买过两次了，但可能选择下一次进行第二次买入花费最小），第i天第二次买入)

// dp[i][4] = max(dp[i - 1][4], dp[i - 1][3] + prices[i])
// max(第i天没有操作沿用i-1天的状态（已经卖出两次了，但可能选择下一次进行第二次卖出钱更多），第i天第二次卖出)

func maxProfit(prices []int) int {
	pl := len(prices)

	if pl == 0 {
		return 0
	}

	// 状态数组初始化
	dp := make([][]int, pl)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 5)
	}

	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][3] = -prices[0]

	for i := 1; i < pl; i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	return dp[pl-1][4]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
