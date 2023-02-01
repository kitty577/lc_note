// 思路：
// 构成回文串：全偶数字符 或 偶数字符加一个奇数字符 -----可借助哈希表统计字符出现次数
// 先撇开奇数的字符作为中心考虑，字符贡献的次数为v/2*2。
// 考虑回文字符串中最多只有一个奇数：根据v/2*2可知,每次ans一直为偶数。可以在发现第一个奇数时,将ans加1，
// 此后再遇到奇数时不再增加

package main

func longestPalindrome(s string) int {
	m := make(map[rune]int)
	for _, b := range s {
		m[b]++
	}
	ans := 0
	for _, v := range m {
		ans += v / 2 * 2
		if v%2 == 1 && ans%2 == 0 {
			ans++
		}
	}
	return ans
}

// 2022.10.3
