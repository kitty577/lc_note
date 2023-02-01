package main

// 使用map，key为string 记录字符串，value为int 记录出现的次数

// 原字符串+字符串的反转字符串 可以组成回文 “lc”和"cl"
// 对称字符串也可以组成回文 "gg"
func longestPalindrome(words []string) int {
	mp := map[string]int{}
	for _, s := range words {
		mp[s]++
	}
	ans := 0
	hadMid := false
	for _, s := range words {
		if isSame(s) { // 碰到对称的字符串，考虑是否要放在回文串中间
			if !hadMid && mp[s]%2 != 0 {
				ans += 2
				hadMid = true
			}
			ans += (mp[s] / 2) * 2 * 2
			mp[s] -= (mp[s] / 2) * 2
		} else {
			pair := min(mp[s], mp[reverse(s)])
			ans += pair * 2 * 2
			mp[s] -= pair
			mp[reverse(s)] -= pair
		}
	}
	return ans
}

func isSame(s string) bool {
	return s[0] == s[1]
}

func reverse(s string) string {
	c := []byte(s)
	c[0], c[1] = c[1], c[0]
	return string(c)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 2022.12.19
