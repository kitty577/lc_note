package main

import "container/heap"

// / topK --- 小顶堆
type myHeap [][2]int // 值--频次

func (m myHeap) Len() int {
	return len(m)
}

func (m myHeap) Less(i, j int) bool {
	return m[i][1] < m[j][1]
}

func (m myHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *myHeap) Push(x interface{}) {
	*m = append(*m, x.([2]int))
}

func (m *myHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	times := make(map[int]int) // key为元素值，value为出现次数
	for _, num := range nums {
		times[num]++
	}

	h := &myHeap{}
	heap.Init(h)

	for key, value := range times {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(h).([2]int)[0])
	}
	return res
}
