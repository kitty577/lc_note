package main

// 要是递增，也即字符串为 全是0 / 全是1 / 一串0后全是1
// 假设下标i为0,1字符的分界点。i之前全是0，i以及i之后全是1
// 用数组count记录，count[i]即表示分界点下标为i时，需要翻转的次数
// 则 count[0]即表示下标为0及之后的字符全是1  === 》 全1字符串  == 》 将字符串中所有的'0'全变为1.
// ==> count[0]= 字符串中0的个数
func minFlipsMonoIncr(s string) int {
	turnAllZero := 0
	for _, v := range s {
		if v == '0' {
			turnAllZero++
		}
	}

	min := turnAllZero
	for _, v := range s {
		if v == '0' {
			turnAllZero--
		} else {
			turnAllZero++
		}

		if turnAllZero < min {
			min = turnAllZero
		}
	}
	
	return min
}
