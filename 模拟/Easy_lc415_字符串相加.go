package main

import "strconv"

// 从末尾开始计算
// 要注意处理进位
func addStrings(num1 string, num2 string) string {
	carry := 0 // 进位
	ans := ""  // 结果

	// 从末尾位开始处理
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || carry != 0; i, j = i-1, j-1 {
		x, y := 0, 0

		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}

		tmp := x + y + carry //计算结果为当前位置的数之和 记得加上进位
		ans = strconv.Itoa(tmp%10) + ans
		carry = tmp / 10 // 新的进位
	}
	return ans
}
