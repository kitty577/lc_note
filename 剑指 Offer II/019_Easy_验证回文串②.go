package main

// 双指针
func validPalindrome(s string) bool {
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

	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		if s[left] != s[right] {
			// 当出现不等的情况，由于最多只能删除一个字符。所以要么删除当前的left,要么删除当前的right
			return isPalindrome(left+1, right) || isPalindrome(left, right-1)
		}
	}

	return true
}
