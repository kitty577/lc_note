package main

// 横向扫描 先求前两个的公共前缀，再用此结果与第三个相比 依次类推
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	ans := strs[0]
	for i := 1; i < len(strs); i++ {
		ans = compare(ans, strs[i])
		if len(ans) == 0 {
			break
		}
	}
	return ans
}

func compare(s1, s2 string) string {
	l := min(len(s1), len(s2))
	index := 0
	for index < l && s1[index] == s2[index] {
		index++
	}
	return s1[:index]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 纵向扫描
// 纵向扫描
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// 2022.12.17
