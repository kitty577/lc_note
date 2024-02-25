package main

// 滑动窗口 + 哈希
// 使用map记录需要的字符及其次数，使用一个int来记录需要的字符个数
// left、end窗口的左右边界，当窗口向后滑动，发现需要的字符时，及时更新map 、 int
// 一旦发现集齐了所有所需要的字符，收缩窗口。更新最小窗口长度以及窗口开始的下标。
// 窗口继续移动 往后查找满足条件的子串
func minWindow(s string, t string) string {
	ls, lt := len(s), len(t)
	if lt == 0 || ls == 0 || lt > ls {
		return ""
	}

	//left窗口左边界，right窗口有边界，start是有效窗口的起始位置
	left, right, start := 0, 0, 0

	// size 记录覆盖子串的最小长度
	size := len(s) + 1
	needMap := make(map[byte]int)

	for _, v := range t {
		needMap[byte(v)]++
	}

	needCount := len(t)
	for right < ls {
		if needMap[s[right]] > 0 {
			needCount--
		}
		needMap[s[right]]--

		// 已经集齐了所有需要的字符，开始收缩窗口
		if needCount == 0 {
			for left < right && needMap[s[left]] < 0 {
				needMap[s[left]]++
				left++
			}
			if right-left+1 < size {
				size = right - left + 1
				start = left
			}

			// 继续向后查找符合的子串
			needCount++
			needMap[s[left]]++
			left++
		}
		right++
	}

	if size == len(s)+1 {
		return ""
	}

	return s[start : start+size]

}
