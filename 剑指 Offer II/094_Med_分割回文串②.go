package main

// 动态规划：count[i]表示截止到i位置，需要分割的次数
// 需要以i下标为结尾的最后一次回文s[j+1:i]次数

func minCut(s string) int {
	n := len(s)
	dp := make([][]bool, n) // dp[i][j]表示s[i...j]是否为回文串
	for i := range dp {
		dp[i] = make([]bool, n)
		for j := range dp[i] {
			dp[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
		}
	}

	count := make([]int, n)
	for i := range count {
		if dp[0][i] {
			continue
		}
		count[i] = n - 1
		for j := 0; j < i; j++ {
			if dp[j+1][i] && count[j]+1 < count[i] {
				count[i] = count[j] + 1
			}
		}
	}
	return count[n-1]
}
