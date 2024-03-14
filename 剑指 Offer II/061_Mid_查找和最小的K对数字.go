package main

import "container/heap"

// topK问题 小顶堆
type Pair struct {
	i int // nums1数组中元素的下标
	j int // nums2数组中元素的下标
}

type myHeap struct {
	pairs []Pair
	nums1 []int
	nums2 []int
}

func (m myHeap) Len() int {
	return len(m.pairs)
}

func (m myHeap) Less(a, b int) bool {
	p1, p2 := m.pairs[a], m.pairs[b]
	return m.nums1[p1.i]+m.nums2[p1.j] < m.nums1[p2.i]+m.nums2[p2.j]
}

func (m myHeap) Swap(a, b int) {
	m.pairs[a], m.pairs[b] = m.pairs[b], m.pairs[a]
}

func (m *myHeap) Push(x any) {
	m.pairs = append(m.pairs, x.(Pair))
}

func (m *myHeap) Pop() any {
	old := m.pairs
	n := len(old)
	top := old[n-1]
	m.pairs = old[:n-1]
	return top
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	len1, len2 := len(nums1), len(nums2)
	res := make([][]int, 0, k)
	m := &myHeap{
		pairs: make([]Pair, 0),
		nums1: nums1,
		nums2: nums2,
	}

	// 由于nums1\nums2都是递增的，因此要最小的几对，都应该是其中一个数组的第一个元素和另外一个数组的其他元素搭配
	for i := 0; i < k && i < len1; i++ {
		m.pairs = append(m.pairs, Pair{i, 0})
	}

	// 在此for循环中，第一次拿到的top的组合是nums1[0],nums2[0]
	// 取出后，往其中添加一个，可能影响结果集的，就是当前的nums1元素搭配nums2的下一个元素
	for m.Len() > 0 && len(res) < k {
		top := heap.Pop(m).(Pair)
		i, j := top.i, top.j
		res = append(res, []int{nums1[i], nums2[j]})
		if j+1 < len2 {
			heap.Push(m, Pair{i, j + 1})
		}
	}

	return res
}
