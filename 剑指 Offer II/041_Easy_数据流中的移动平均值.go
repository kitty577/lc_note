package main

// 数组存储队列中的数
type MovingAverage struct {
	nums      []int
	sum, size int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{size: size}
}

func (this *MovingAverage) Next(val int) float64 {
	// 确保数组中的只有size个数
	if len(this.nums) == this.size {
		this.sum = this.sum - this.nums[0]
		this.nums = this.nums[1:]
	}

	this.sum += val
	this.nums = append(this.nums, val)
	return float64(this.sum) / float64(len(this.nums))
}
