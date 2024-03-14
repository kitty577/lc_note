package main

// 桶
// 按照元素的大小分桶，对于一个元素x,与之可能满足差值条件的元素大小为[x-valueDiff,x+valueDiff]
// 设定桶的大小为 valueDiff+1
// 如果两个元素同属一个桶，那么这两个元素必然符合条件。
// 如果两个元素属于相邻桶，那么我们需要校验这两个元素是否差值不超过 valueDiff
// 如果两个元素既不属于同一个桶，也不属于相邻桶，那么这两个元素必然不符合条件。
func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	m := make(map[int]int) // key为桶id,value为元素值
	for i, v := range nums {
		bid := getID(v, valueDiff+1)
		// 桶里有数
		if _, has := m[bid]; has {
			return true
		}
		// 桶里没有，找左右相邻的桶
		if y, has := m[bid+1]; has && abs(v-y) <= valueDiff {
			return true
		}
		if y, has := m[bid-1]; has && abs(v-y) <= valueDiff {
			return true
		}
		m[bid] = v

		// 滑动窗口,移除窗口外的
		if i >= indexDiff {
			delete(m, getID(nums[i-indexDiff], valueDiff+1))
		}
	}
	return false
}

// 根据值大小和桶宽度来计算相应的桶ID
func getID(x, w int) int {
	if x >= 0 {
		return x / w
	}
	// 当输入的值为负数时，考虑向下取整的情况。 整数除法的结果会向0取整，而不是向下取整。
	// 这意味着对于负数，除法结果会向着0靠拢，而不是向负无穷靠拢。

	// x+1 保证结果向下取整； 再减去1，是因为索引是从0开始，而不是从1开始
	return ((x + 1) / w) - 1
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
