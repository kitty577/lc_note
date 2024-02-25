package main

// 栈
func asteroidCollision(asteroids []int) []int {
	stack := []int{}

	for _, a := range asteroids {
		alive := true // 标记当前处理的星球a是否存活

		for alive && len(stack) > 0 && stack[len(stack)-1] > 0 && a < 0 {
			alive = abs(a) > abs(stack[len(stack)-1])
			if abs(stack[len(stack)-1]) <= abs(a) {
				stack = stack[:len(stack)-1]
			}
		}

		if alive {
			stack = append(stack, a)
		}
	}
	return stack
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
