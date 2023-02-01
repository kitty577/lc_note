package main

import "strconv"

//可以通过模拟「竖式乘法」的方法计算乘积。从右往左遍历乘数，将乘数的每一位与被乘数相乘得到对应的结果，再将每次得到的结果累加。

// TODO:除了最低位以外，其余的每一位的运算结果都需要补 0

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	ans := "0"
	m, n := len(num1), len(num2)
	for i := n - 1; i >= 0; i-- {
		carry := 0
		cur := ""

		for j := n - 1; j > i; j-- {
			cur += "0" // 低位补0
		}

		y := int(num2[i] - '0')
		for j := m - 1; j >= 0; j-- {
			x := int(num1[j] - '0')
			tmp := x*y + carry
			cur = strconv.Itoa(tmp%10) + cur
			carry = tmp / 10
		}

		for carry != 0 {
			cur = strconv.Itoa(carry%10) + cur
			carry = carry / 10
		}

		ans = addStrings(ans, cur)
	}

	return ans
}

func addStrings(num1 string, num2 string) string {
	ans := ""
	carry := 0

	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || carry != 0; i, j = i-1, j-1 {
		x, y := 0, 0

		if i >= 0 {
			x = int(num1[i] - '0')
		}

		if j >= 0 {
			y = int(num2[j] - '0')
		}

		tmp := x + y + carry
		ans = strconv.Itoa(tmp%10) + ans
		carry = tmp / 10
	}
	return ans
}
