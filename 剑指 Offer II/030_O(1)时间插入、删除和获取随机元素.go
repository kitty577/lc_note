package main

import "math/rand"

type RandomizedSet struct {
	nums     []int
	indexMap map[int]int // key为元素值，value为该元素在数组中的下标
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		nums:     []int{},
		indexMap: make(map[int]int),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.indexMap[val]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	this.indexMap[val] = len(this.nums) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.indexMap[val]; !ok {
		return false
	}
	index := this.indexMap[val]
	l := len(this.nums)
	this.nums[index] = this.nums[l-1]
	this.indexMap[this.nums[index]] = index
	this.nums = this.nums[:l-1]
	delete(this.indexMap, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
