// 思路：
// 首先，将原始字符串做初步处理（删除掉空格、破折号）---可借助strings库，strings.ReplaceAll
// 再者，由于要求将数组从左到右 每 3 个一组 分块，直到 剩下 4 个或更少数字。剩下的数字将按下述规定再分块：
// 		2 个数字：单个含 2 个数字的块。
// 		3 个数字：单个含 3 个数字的块。
// 		4 个数字：两个分别含 2 个数字的块
//     可通过截取slice切片实现。
// 最后，将分好组的数字块用破折号连接起来---strings.Join

package main

import "strings"

func reformatNumber(number string) string {
	s := strings.ReplaceAll(number, " ", "")
	s = strings.ReplaceAll(s, "-", "")
	ans := []string{}
	i := 0
	for ; i+4 < len(s); i += 3 {
		ans = append(ans, s[i:i+3])
	}
	s = s[i:]
	if len(s) < 4 {
		ans = append(ans, s)
	} else {
		ans = append(ans, s[:2], s[2:])
	}
	return strings.Join(ans, "-")
}

// 2022.10.1
