package main

// 双指针
// 子串开始的下标和结束的下标任意，因此从头开始遍历
func countSubstrings(s string) int {
	ans, n := 0, len(s)
	isPalindrome := func(i, j int) bool {
		for i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}

	for start := 0; start < n; start++ {
		for end := start; end < n; end++ {
			if s[start] == s[end] {
				if isPalindrome(start, end) {
					ans++
				}
			}
		}
	}
	return ans
}
