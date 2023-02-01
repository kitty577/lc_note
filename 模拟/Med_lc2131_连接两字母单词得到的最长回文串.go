package main

// 与409最长回文串相似
// 想到了使用map，key为string 记录字符串，value为int 记录出现的次数
func longestPalindrome(words []string) int {
	m := map[string]int{}
	for _, str := range words {
		m[str]++
	}

	res := 0
	hasMid := false
	for _, word := range words {
		if isSame(word) {
			if !hasMid && m[word]%2 == 1 {
				res += len(word)
				hasMid = true
			}
			res += (m[word] / 2) * 4
			m[word] -= (m[word] / 2) * 2
		} else {
			cnt := min(m[reverse(word)], m[word])
			res += cnt * 4
			m[reverse(word)] -= cnt
			m[word] -= cnt
		}
	}
	return res
}

func reverse(s string) string {
	b := []byte(s)
	b[0], b[1] = b[1], b[0]
	return string(b)
}

func isSame(s string) bool {
	return s[0] == s[1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
