package main

// 找到第一个大（小）的元素 ==》典型单调栈
// 这道题中 找到下一个更大的数 也就是说数本身是要递增顺序的
// 由于栈是先进后出。要保证出栈的顺序的递增   那么入栈则要递减
func dailyTemperatures(num []int) []int {
    ans := make([]int, len(num)) // 结果
    stack := []int{}  // 模拟栈 存下标
    for i, v := range num {
        // 栈不空，且当前遍历元素 v 破坏了栈的单调性
        for len(stack) != 0 && v > num[stack[len(stack)-1]] {
            // pop
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            ans[top] = i - top
        }
        stack = append(stack, i)
    }
    return ans
}
