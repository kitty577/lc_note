package main

func partition(s string) [][]string {
	res := [][]string{}
	tmp := []string{}
	find(s, 0, tmp, &res)
	return res
}

func find(s string, start int, tmpString []string, res *[][]string) {
	// 到达字符串末尾
	if start == len(s) {
		tmp := make([]string, len(tmpString))
		copy(tmp, tmpString)
		*res = append(*res, tmp)
		return
	}

	for i := start; i < len(s); i++ {
		if isPalindrome(s, start, i) {
			tmpString = append(tmpString, s[start:i+1])
		} else {
			continue
		}

		find(s, i+1, tmpString, res)

		// 还原现场
		tmpString = tmpString[:len(tmpString)-1]
	}

}

func isPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}

		start++
		end--
	}

	return true
}
