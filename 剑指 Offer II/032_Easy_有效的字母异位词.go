package main

// 哈希
func isAnagram(s string, t string) bool {
	ls, lt := len(s), len(t)
	if ls != lt {
		return false
	}
	if s == t {
		return false
	}

	cs, ct := [26]int{}, [26]int{}
	for _, v := range s {
		cs[v-'a']++
	}
	for _, v := range t {
		ct[v-'a']++
	}
	return cs == ct
}
