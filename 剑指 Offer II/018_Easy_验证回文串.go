package main

import "strings"

// 双指针
// 首先对s处理下，去除空格、统一为小写
func isPalindrome(s string) bool {
	var ns string
	for i := 0; i < len(s); i++ {
		if isValid(s[i]) {
			ns += string(s[i])
		}
	}

	ns = strings.ToLower(ns)
	l := len(ns)
	for i := 0; i < l/2; i++ {
		if ns[i] != ns[l-i-1] {
			return false
		}
	}
	return true
}

func isValid(v byte) bool {
	return (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9')
}
