// 解读下题目内容，题意中 满足同构：一个字符映射到另外一个字符，s 和 t 每个位置上的字符是否都一一对应，即 s 的任意一个字符被 t 中唯一的字符对应，同时 t 的任意一个字符被 s 中唯一的字符对应。
// 想到哈希表
// 可以维护两个哈希表：m1的key为s中的字符，value为t中对应的字符。记录s到t的映射关系
//                   m2的key为t中的字符，value为s中对应的字符。记录t到s的映射关系
// 遍历字符串，不断更新m1,m2
package main

func isIsomorphic(s string, t string) bool {
	s2tM := make(map[byte]byte)
	t2sM := make(map[byte]byte)
	for i := range s {
		si, ti := s[i], t[i]
		// 此处要注意，s到t、t到s都是唯一映射！
		if s2tM[si] > 0 && s2tM[si] != ti || t2sM[ti] > 0 && t2sM[ti] != si {
			return false
		}
		s2tM[si] = ti
		t2sM[ti] = si
	}
	return true
}

// 2022.10.1
