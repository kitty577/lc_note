package main

// 前缀树
type tire struct {
	children [26]*tire
	val      int // 当前节点的权重。
	// 根据题意，对于一个字典key,从0开始到任意位置的截断的前缀权重都为val.
	// 如，对于“apple”来说，“a”、"ap"、“app”....等前缀树上的节点的权重都为apple对应的val
}

type MapSum struct {
	t   *tire
	cnt map[string]int
}

func Constructor() MapSum {
	return MapSum{
		t:   &tire{},
		cnt: make(map[string]int),
	}
}

func (this *MapSum) Insert(key string, val int) {
	delta := val
	// 如果cnt维护的map中已经存在key,则说明key已经建立前缀树，需要更新节点权重
	// 如果不存在，则构建前缀树
	if _, ok := this.cnt[key]; ok {
		delta = val - this.cnt[key]
	}
	this.cnt[key] = val
	cur := this.t
	for _, ch := range key {
		ch = ch - 'a'
		if cur.children[ch] == nil {
			cur.children[ch] = &tire{}
		}
		cur = cur.children[ch]
		cur.val += delta
	}
}

func (this *MapSum) Sum(prefix string) int {
	cur := this.t
	for _, p := range prefix {
		p = p - 'a'
		if cur.children[p] == nil {
			return 0
		}
		cur = cur.children[p]
	}
	return cur.val
}
