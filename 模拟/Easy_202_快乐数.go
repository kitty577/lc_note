// 由题意可知 重复这个过程 要么结果是1、要么是无限循环
// ===》转化为是否含有环==》快慢指针

// 慢跳一步 快跳两步

package main

func isHappy(n int) bool {
	slow, fast := n, step(n)
	for fast != 1 && slow != fast {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

// 2022.10.17
