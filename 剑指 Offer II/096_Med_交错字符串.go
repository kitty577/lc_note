package main

// 动态规划
// dp[i][j]表示s1[0...i]与s2[0...j]能否满足交错字符串

func isInterleave(s1 string, s2 string, s3 string) bool {
	len1, len2, len3 := len(s1), len(s2), len(s3)
	if len1+len2 != len3 {
		return false
	}
	dp := make([][]bool, len1+1)
	for i := range dp {
		dp[i] = make([]bool, len2+1)
	}
	dp[0][0] = true
	for i := range s1 {
		for j := range s2 {
			tail := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || dp[i-1][j] && s3[tail] == s1[i-1]
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || dp[i][j-1] && s3[tail] == s2[j-1]
			}
		}
	}

	return dp[len1][len2]
}
