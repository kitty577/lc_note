package main

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	ans := make([][]string, 0)
	m := map[[26]int][]string{}
	for _, s := range strs {
		cnt := [26]int{}
		for _, v := range s {
			cnt[v-'a']++
		}
		m[cnt] = append(m[cnt], s)
	}
	for _, vs := range m {
		ans = append(ans, vs)
	}
	return ans
}
