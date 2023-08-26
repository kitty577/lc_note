package main

// 如果 i 为偶数，即二进制的末位为0；那么i 的1的个数与 i>>1 中1的个数相同
// 如果 i 为奇数，那么i中含1的个数也就是 i>>1中1的个数+1

// 判断一个二进制数的奇偶，也就是判断二进制数的末位上数字是否是1  ==》 与0x1做与操作 就可以得出
func countBits(n int) []int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = dp[i>>1] + i&0x1
	}
	return dp
}
