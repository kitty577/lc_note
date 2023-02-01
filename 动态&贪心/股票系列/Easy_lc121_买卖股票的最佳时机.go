// 思路：
// 最大利润=最高价格-最低价格
package main

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0
	for _, dayPrice := range prices {
		maxProfit = max(maxProfit, dayPrice-minPrice)
		minPrice = min(minPrice, dayPrice)
	}
	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// 2022.10.3
