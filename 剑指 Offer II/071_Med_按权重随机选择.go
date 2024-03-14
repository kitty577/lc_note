package main

import "math/rand"

// 前缀和
// 设总权重之和为total
// w=[3,1,2,4]，权重之和为3+1+2+4=10
// 前缀和[3,4,6,10]
// 随机产生一个随机数、找到这个随机数在前缀和数组中的下标位置即可
type Solution struct {
	pre []int
}

func Constructor(w []int) Solution {
	length := len(w)
	if length <= 1 {
		return Solution{pre: w}
	}
	for i := 1; i < length; i++ {
		w[i] += w[i-1]
	}
	return Solution{pre: w}
}

// rand.Intn(n) 代表从[0,n）中生成一个随机数.
func (this *Solution) PickIndex() int {
	x := rand.Intn(this.pre[len(this.pre)-1]) + 1 // [1,n]
	return search(this.pre, x)
}

func search(arr []int, n int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == n {
			return mid
		} else if arr[mid] > n {
			right = mid - 1
		} else if arr[mid] < n {
			left = mid + 1
		}
	}
	return left
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
