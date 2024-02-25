package main

// 下一个更大的元素
// 单调栈
func dailyTemperatures(temperatures []int) []int {
	stack := []int{} // 栈 存储元素下标
	ans := make([]int, len(temperatures))
	for i, t := range temperatures {
		for len(stack) != 0 && temperatures[stack[len(stack)-1]] < t {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[top] = i - top
		}
		stack = append(stack, i)
	}
	return ans
}
