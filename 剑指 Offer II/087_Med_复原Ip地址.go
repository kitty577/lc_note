package main

import "strconv"

const (
	spilt       = "."
	partsLength = 4
)

func restoreIpAddresses(s string) []string {
	res := []string{}  // ip的数组
	path := []string{} // 一个ip被分割成四段
	find(s, 0, &path, &res)
	return res
}

func find(s string, start int, path *[]string, res *[]string) {
	// 分割完成
	if len(*path) == partsLength && start == len(s) {
		tmpIp := (*path)[0] + spilt + (*path)[1] + spilt + (*path)[2] + spilt + (*path)[3]
		*res = append(*res, tmpIp)
		return
	}

	for i := start; i < len(s); i++ {
		if len(*path) < 4 && isValid(s, start, i) {
			*path = append(*path, s[start:i+1])
		} else {
			continue
		}

		find(s, i+1, path, res)

		// 还原
		*path = (*path)[:len(*path)-1]
	}
}

// 判断子段是否合法
func isValid(s string, start, end int) bool {
	if end-start+1 < 0 || end-start+1 > 3 || end > start && s[start] == '0' {
		return false
	}
	checkInt, _ := strconv.Atoi(s[start : end+1])
	if checkInt > 255 {
		return false
	}

	return true
}
