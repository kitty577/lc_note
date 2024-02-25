package main

// 方法一：首先想到的是直接遍历。一个map记录order的字母顺序，然后对words中的字符串，与前一个字符串比较（各个字母进行比较）
// 方法二：已知是英文小写字母，所以可以和原本英文字母顺序建立映射关系，将words中的字符串中的字母转化为映射后的字母，直接比较
func isAlienSorted(words []string, order string) bool {
	// 建立映射关系
	orderB := []byte(order)
	m := map[byte]byte{}
	for i, v := range orderB {
		m[v] = byte(i + 'a')
	}
	var pre string
	for _, word := range words {
		w2b := []byte(word)
		for i, v := range w2b {
			w2b[i] = m[v]
		}
		b2s := string(w2b)
		if pre > b2s {
			return false
		}
		pre = b2s
	}
	return true
}
