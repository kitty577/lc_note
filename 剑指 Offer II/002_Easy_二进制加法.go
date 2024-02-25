package main

import "strconv"

// 从低位开始往高位计算
// 对两数的同一位上而言：
// 如果两数上都是1 则相加要进位====》 用 与 判断是否需要进位
// 处理完进位后  其他位上可以用 异或 判断相加后该位上的数
func addBinary(a string, b string) string {
	ans := ""
	// 从低位开始计算
	i := len(a) - 1
	j := len(b) - 1

	// 进位
	carry := 0

	for i >= 0 || j >= 0 {
		digitA := 0
		digitB := 0
		if i >= 0 {
			digitA = int(a[i] - '0')
			i--
		}
		if j >= 0 {
			digitB = int(b[j] - '0')
			j--
		}

		sum := digitA + digitB + carry
		if sum >= 2 {
			carry = 1
			sum = sum - 2
		} else {
			carry = 0
		}

		ans = strconv.Itoa(sum) + ans
	}

	// 处理完所有位 还有进位
	if carry != 0 {
		ans = strconv.Itoa(carry) + ans
	}

	return ans
}
