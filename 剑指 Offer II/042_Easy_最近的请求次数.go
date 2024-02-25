package main

type RecentCounter struct {
	times []int
}

func Constructor() RecentCounter {
	return RecentCounter{times: []int{}}
}

func (this *RecentCounter) Ping(t int) int {
	this.times = append(this.times, t)
	for this.times[0] < t-3000 {
		this.times = this.times[1:]
	}
	return len(this.times)
}
