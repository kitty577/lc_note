package main

// 乘除优先于加减。
// 使用栈。对于加减号前后的数字，直接压入栈。对于乘除号后的数组，取出栈顶元素与之运算，并将运算结果代替栈顶元素压入栈

// 最后 栈中所有元素相加，即为所求

func calculate(s string) (ans int) {
	stack := []int{}

	preSign := '+' // 数字前一个运算符

	num := 0

	for i, ch := range s {
		isDigit := '0' <= ch && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0')
		}
		if !isDigit && ch != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] = stack[len(stack)-1] * num
			case '/':
				stack[len(stack)-1] = stack[len(stack)-1] / num
			}

			preSign = ch
			num = 0
		}
	}
	for _, v := range stack {
		ans += v
	}
	return
}
