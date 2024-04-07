package main

// 动态规划，dp[len1+1][len2+1]// ps：加1 防止数组越界，不用为len=0/len==1单独写逻辑
// dp[i][j]表示text1[0:i]和text2[0:j]的最长公共子序列的长度
// i=0  也即text1[0:0] 空字符串，因此 dp[0][j]=0
// 同理 dp[i][0]=0
// i>0&&j>0时，text1[i-1]==text2[j-1]  dp[i][j]=dp[i-1][j-1]+1
//
//	text1[i-1]!=text2[j-1]  dp[i][j]=max(dp[i-1][j],dp[i][j-1])
func longestCommonSubsequence(text1 string, text2 string) int {
	len1, len2 := len(text1), len(text2)
	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}

	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[len1][len2]
}
