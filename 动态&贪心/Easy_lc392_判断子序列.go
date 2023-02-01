package main

// 思路：
// 方法一：双指针
//         子序列，由于是s中的全部元素都在t中，相对位置不变；因此只需要顺序遍历，记录下匹配到的长度
// func isSubsequence(s string, t string) bool {
//     n,m:=len(s),len(t)
//     i,j:=0,0
//     for i<n&&j<m{
//         if s[i]==t[j]{
//             i++
//         }
//         j++
//     }
//     return i==n
// }

// 方法二：动态规划
//          dp[i][j]记录匹配到长度。以下标i-1结尾的字符串s,和以下标j-1结尾的字符串t,相同子序列的长度为dp[i][j]
func isSubsequence(s string, t string) bool {
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(t)+1)
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(s)][len(t)] == len(s)

}

// 2022.10.8
