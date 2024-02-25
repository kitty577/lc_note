package main

// 单调栈
// 木桶原理。遍历到高度低于栈顶时，则说明要更新栈
func largestRectangleArea(heights []int) int {
	heights = append(heights, 0) // 末尾添0，保证能够更新栈，不会漏掉最末尾的一次结果
	stack := []tangle{}
	ans := 0
	for _, h := range heights {
		accumulateWidth := 0
		for len(stack) != 0 && stack[len(stack)-1].h >= h {
			top := stack[len(stack)-1]
			accumulateWidth += top.w
			ans = max(ans, accumulateWidth*top.h)
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, tangle{accumulateWidth + 1, h})
	}
	return ans
}

type tangle struct {
	w int
	h int
}
