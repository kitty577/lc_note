package main

// 滑动窗口+哈希

func lengthOfLongestSubstring(s string) int {
	ls := len(s)
	if ls < 2 {
		return ls
	}
	visited := make(map[byte]bool)
	start, end, ans := 0, 0, 0
	for end < ls {
		if visited[s[end]] {
			visited[s[start]] = false
			start++
		} else {
			visited[s[end]] = true
			ans = max(ans, end-start+1)
			end++
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
