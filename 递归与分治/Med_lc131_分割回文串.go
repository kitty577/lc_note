package main

func partition(s string) [][]string {
	// 回溯法，分割回文子串
	tmpString := []string{}
	res := [][]string{}
	backtrack(s, tmpString, 0, &res)
	return res
}

func backtrack(s string, tmpString []string, startIndex int, res *[][]string) {
	// 到达字符串末尾
	if startIndex == len(s) {
		tmp := make([]string, len(tmpString))
		copy(tmp, tmpString)
		*res = append(*res, tmp)
		return
	}
	// 回溯单层逻辑
	for i := startIndex; i < len(s); i++ {
		// 如果回文，则处理节点，否则继续遍历
		if isPalindrome(s, startIndex, i) {
			tmpString = append(tmpString, s[startIndex:i+1])
		} else {
			continue
		}
		// 递归
		backtrack(s, tmpString, i+1, res)
		// 回溯，撤销节点
		tmpString = tmpString[:len(tmpString)-1]
	}
}

// 回文判断
func isPalindrome(s string, start, end int) bool {
	left, right := start, end
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 2022.12.9
