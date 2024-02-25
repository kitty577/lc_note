package main

// 由于LRUCache 有容量限制，因此使用capacity、size 记录容量和当前使用量。 使用map记录缓存中是否有该元素
// 另外 由于需要记录最近最少使用。 因此想到双向链表
type node struct {
	key, value int
	pre, next  *node
}
type LRUCache struct {
	capacity   int
	size       int
	m          map[int]*node
	head, tail *node // 这里的头尾节点，都只是为了方便对链表的操作，其实都是空节点，相当于dummyHead
}

func InitNode(key, val int) *node {
	return &node{
		key:   key,
		value: val,
	}
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		size:     0,
		m:        make(map[int]*node),
		head:     InitNode(0, 0),
		tail:     InitNode(0, 0),
	}
	lru.head.next = lru.tail
	lru.tail.pre = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.m[key]; !ok {
		return -1
	}
	node := this.m[key]
	// 使用到该节点，需要把该节点挪到链表头部
	this.MoveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.m[key]; ok {
		node := this.m[key]
		// 关键字存在，则修改节点值，并挪到链表头部
		node.value = value
		this.MoveToHead(node)

	} else {
		// 不存在，则新建node
		new := InitNode(key, value)
		// 把该结点添加到头部，判断容量、更改size、更新map
		this.m[key] = new
		this.AddToHead(new)
		this.size++
		if this.size > this.capacity {
			removeNode := this.RemoveTail()
			delete(this.m, removeNode.key)
			this.size--
		}
	}
}

// 将节点添加到头部
func (this *LRUCache) AddToHead(node *node) {
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
}

// 将节点移动到头部
func (this *LRUCache) MoveToHead(node *node) {
	this.Remove(node)
	this.AddToHead(node)
}

// 将末尾节点删除
func (this *LRUCache) RemoveTail() *node {
	node := this.tail.pre
	this.Remove(node)
	return node
}

// 删除节点
func (this *LRUCache) Remove(node *node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
