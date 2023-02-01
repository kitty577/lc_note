package main

import "strconv"

// 这题和131.分割回文串 很像
func restoreIpAddresses(s string) []string {
	res := []string{}
	path := []string{} // 记录单次分割结果
	find(s, path, 0, &res)
	return res

}

func find(s string, path []string, startIndex int, res *[]string) {
	// 首先 递归终止条件：边界条件+分成四段
	if startIndex == len(s) && len(path) == 4 {
		tmpIp := path[0] + "." + path[1] + "." + path[2] + "." + path[3]
		*res = append(*res, tmpIp)
		return
	}

	// 单层回溯
	for i := startIndex; i < len(s); i++ {
		// 一段的数字个数最多为3
		if i-startIndex+1 <= 3 && len(path) < 4 && isValid(s, startIndex, i) {
			path = append(path, s[startIndex:i+1])
			find(s, path, i+1, res)
		} else {
			return
		}

		// 还原现场
		path = path[:len(path)-1]
	}
}

// 判断子段是否合法
func isValid(s string, start, end int) bool {
	// 不含前导0 (strat、end截取的长度大于1)
	if end > start && s[start] == '0' {
		return false
	}

	checkInt, _ := strconv.Atoi(s[start : end+1])
	if checkInt > 255 {
		return false
	}

	return true
}

// 2022.12.10
