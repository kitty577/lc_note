package main

// 循环 也就是遍历两遍数组
func nextGreaterElements(nums []int) []int {
	length := len(nums)
	ans := make([]int, length) // 存放最终结果。默认值为-1
	for i := 0; i < length; i++ {
		ans[i] = -1
	}

	stack := []int{} // 模拟栈，存放下一个大的元素的下标
	for i := 0; i < length*2; i++ {
		for len(stack) != 0 && nums[i%length] > nums[stack[len(stack)-1]] {
			// pop
			ans[stack[len(stack)-1]] = nums[i%length]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%length)
	}
	return ans
}
