package main

// 线段树
// 线段树的每个节点都代表一个区间。
// 线段树具有唯一的根节点，代表的区间是整个统计范围，比如 [1, n]。
// 线段树的每个叶子节点都代表一个长度为 1 的单位区间 [x, x]。
// 对于每个内部节点 [left, right]，它的左子节点是 [left, mid]，右子节点是 [mid + 1, right]。其中 mid = (left + right) / 2（向下取整）

// 节点编号为n,则其左孩子节点2*n+1, 右孩子节点2*n+2
type MyCalendar struct {
	// 线段树中每个节点表示一个区间
	tree map[int]bool // 记录 节点编号idx(也就是一个区间) 中是否有部分被预订
	lazy map[int]bool // 标记idx区间是否被预订

}

func Constructor() MyCalendar {
	return MyCalendar{map[int]bool{}, map[int]bool{}}
}

func (c MyCalendar) query(start, end, l, r, idx int) bool {
	if r < start || end < l { // 完全处于[l, r]左边 或 右边
		return false
	}
	if c.lazy[idx] { // 如果该区间已被预订，则直接返回
		return true
	}
	if start <= l && r <= end { // 所求的start,end甚至完全囊括[l, r]区间
		return c.tree[idx]
	}
	mid := (l + r) >> 1
	return c.query(start, end, l, mid, 2*idx+1) ||
		c.query(start, end, mid+1, r, 2*idx+2)
}

func (c MyCalendar) update(start, end, l, r, idx int) {
	if r < start || end < l { // 完全处于[l, r]左边 或 右边
		return
	}
	if start <= l && r <= end { // 所求的start,end甚至完全囊括[l, r]区间
		c.tree[idx] = true
		c.lazy[idx] = true
	} else { // 所求的start,end区间 部分 或完全处于[l, r] 中，需要更新具体线段的标记
		mid := (l + r) >> 1
		c.update(start, end, l, mid, 2*idx+1)
		c.update(start, end, mid+1, r, 2*idx+2)
		c.tree[idx] = true
		if c.lazy[2*idx+1] && c.lazy[2*idx+2] {
			c.lazy[idx] = true
		}
	}
}

func (c MyCalendar) Book(start, end int) bool {
	if c.query(start, end-1, 0, 1e9, 0) {
		return false
	}
	c.update(start, end-1, 0, 1e9, 0)
	return true
}
