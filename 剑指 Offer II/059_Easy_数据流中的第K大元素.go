package main

import "container/heap"

// topK问题： 使用最小堆来解决
// 大顶堆： 父节点都大于左右孩子结点
// 小顶堆： 父节点都小于左右孩子结点
type PriorQ []int

type KthLargest struct {
	PQ   *PriorQ
	Size int
}

func Constructor(k int, nums []int) KthLargest {
	p := &PriorQ{}
	kth := KthLargest{
		p,
		k,
	}
	for _, n := range nums {
		kth.Add(n)
	}
	return kth
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.PQ, val)
	if this.PQ.Len() > this.Size {
		heap.Pop(this.PQ)
	}
	return (*this.PQ)[0]

}

// 为PriorQ实现heap接口
func (p PriorQ) Len() int {
	return len(p)
}

func (p PriorQ) Less(i int, j int) bool {
	return p[i] < p[j]
}

func (p PriorQ) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorQ) Push(x any) {
	v := x.(int)
	*p = append(*p, v)
}

func (p *PriorQ) Pop() any {
	old := *p
	n := len(old)
	top := old[n-1]
	*p = old[:n-1]
	return top
}
