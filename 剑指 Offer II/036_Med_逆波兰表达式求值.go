package main

import "strconv"

// 栈
// 遇到数字，入栈；遇到符号，弹出栈顶两元素，进行运算，将计算后的结果压入栈顶

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, s := range tokens {
		val, err := strconv.Atoi(s)
		if err == nil {
			// 数字 --> 入栈
			stack = append(stack, val)
		} else {
			// 符号 --> 弹出栈顶两元素，运算
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch s {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}
