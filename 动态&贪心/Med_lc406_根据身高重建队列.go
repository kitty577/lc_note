package main

import "sort"

// 先按身高排序
// 当放入第I个人时，第0~i-1个人已经排好了位置，且无论他们怎么站都对第i个人没有影响，因为他们比i矮
// 而第i+1以及之后的人还没有放入位置，且他们的站位对i有影响

// 可以先建立个空队列，每次将一个人放入队列中时，会将一个空位置变成满位置。那么当处理第i个人时，我们需要给他找一个空位置，且它的前面还有要ki个空位置，用于安排后面更高的人。换言之，第i个人的位置 其实就是从左到右的第ki+1个空位置
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] < people[j][0] || people[i][0] == people[j][0] && people[i][1] > people[j][1]
	})

	ans := make([][]int, len(people))

	for _, v := range people {
		space := v[1] + 1
		for i := range ans {
			if ans[i] == nil {
				space--
				if space == 0 {
					ans[i] = v
					break
				}
			}
		}
	}

	return ans
}
