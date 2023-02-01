package main

// 由于爆炸发生在相邻的两颗行星之间 因此想到栈

// 当我们遍历到行星aster 时，使用变量alive 记录行星aster 是否还存在（即未爆炸）

// 当行星存在且是负的时，如果栈顶是正，那么爆炸发生：
//      如果栈顶正的更多，那么负的爆炸。这个行星不能存活
//      如果这个行星负的更多，那么这个行星存活，正的爆炸，弹出栈

func asteroidCollision(asteroids []int) []int {
	res := []int{}

	for _, aster := range asteroids {
		alive := true
		for alive && aster < 0 && len(res) > 0 && res[len(res)-1] > 0 { // 发生爆炸
			alive = res[len(res)-1] < abs(aster)
			if res[len(res)-1] <= abs(aster) { // ﹣的更大
				res = res[:len(res)-1]
			}
		}
		if alive {
			res = append(res, aster)
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

// 2023.01.05
