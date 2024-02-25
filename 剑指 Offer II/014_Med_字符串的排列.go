package main

// 滑动窗口+哈希

func checkInclusion(s1 string, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	if l1 > l2 {
		return false
	}
	c1, c2 := [26]int{}, [26]int{}
	for i, v := range s1 {
		c1[v-'a']++
		c2[s2[i]-'a']++
	}
	if c1 == c2 {
		return true
	}

	for i := 1; i <= l2-l1; i++ {
		c2[s2[i-1]-'a']--
		c2[s2[i+l1-1]-'a']++
		if c1 == c2 {
			return true
		}
	}
	return false
}
