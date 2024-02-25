package main

// 滑动窗口+哈希
func findAnagrams(s string, p string) []int {
	ans := []int{}
	ls, lp := len(s), len(p)
	if lp > ls {
		return ans
	}
	cs, cp := [26]int{}, [26]int{}
	for i, v := range p {
		cp[v-'a']++
		cs[s[i]-'a']++
	}
	if cs == cp {
		ans = append(ans, 0)
	}

	for i := 1; i <= ls-lp; i++ {
		cs[s[i-1]-'a']--
		cs[s[i+lp-1]-'a']++
		if cs == cp {
			ans = append(ans, i)
		}
	}
	return ans
}
