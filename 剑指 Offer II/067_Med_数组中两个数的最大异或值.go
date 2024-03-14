package main

// 前缀树
// 异或： 只有两个二进制位不同时，结果才为1。 因此要尽可能找到两个很多位都不同的数，且先考虑高位。
// 构建前缀树，将所有数都插入到前缀树中，类似于字符串构建是从头到尾按字符，这里从高位到低位插入数的二进制位

const highBit = 30 // 根据题意，值为整数，不用考虑负数。因此只要考虑31位，那么右移最多三十位

type tire struct {
	left, right *tire // left为二进制位为0的分支，right为二进制位1的分支
}

func (t *tire) add(num int) {
	cur := t
	for i := highBit; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			if cur.left == nil {
				cur.left = &tire{}
			}
			cur = cur.left
		} else {
			if cur.right == nil {
				cur.right = &tire{}
			}
			cur = cur.right
		}
	}
}

func (t *tire) check(num int) int {
	var ans int
	cur := t
	for i := highBit; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 { // ---- 当前二进制位为0，那么要使异或值大，则需要找二进制为1的。
			if cur.right != nil {
				cur = cur.right
				ans = ans*2 + 1
			} else {
				cur = cur.left
				ans = ans * 2
			}
		} else { // ---- 当前二进制位为1
			if cur.left != nil {
				cur = cur.left
				ans = ans*2 + 1
			} else {
				cur = cur.right
				ans = ans * 2
			}
		}
	}
	return ans
}
func findMaximumXOR(nums []int) int {
	var res int
	root := &tire{}
	for i := 1; i < len(nums); i++ {
		// 可以同时add和check,是因为A异或B 与B异或A 答案是一样的，不会存在结果遗漏掉当前未add的数
		root.add(nums[i-1])
		res = max(res, root.check(nums[i]))
	}
	return res
}
