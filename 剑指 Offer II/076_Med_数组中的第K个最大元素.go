package main

import "container/heap"

// topK问题: 小顶堆

type heapQ []int

// heap接口实现
func (q heapQ) Len() int {
	return len(q)
}

func (q heapQ) Less(i, j int) bool {
	return q[i] < q[j]
}

func (q heapQ) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *heapQ) Push(x any) {
	*q = append(*q, x.(int))
}

func (q *heapQ) Pop() any {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[:n-1]
	return x
}

// 往小顶堆里加入数据
func (q *heapQ) Add(val, k int) {
	// 如果堆的长度小于k， 或者 会破坏大小顺序， 插入数据到堆中
	if q.Len() < k || val > (*q)[0] {
		heap.Push(q, val)
	}
	// 保证堆的长度不超过K
	if q.Len() > k {
		heap.Pop(q)
	}
}

func findKthLargest(nums []int, k int) int {
	// 建立堆
	var q heapQ
	for _, v := range nums {
		q.Add(v, k)
	}
	return q[0]
}
