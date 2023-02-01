package main

// 双向链表+哈希
// 双向链表的顺序按照节点的使用顺序来：首为最近使用的 尾为最久未使用的
// get操作为：读取哈希表中的值、将改节点的位置移到链表头部（更新）
// put操作为：如果缓存中已经存在该值，则更新值，更新节点在链表中的位置
//            如果缓存中该值不存在，则如果缓存中还有位置，则插入到头部；
//                                    如果缓存中没有位置，则删去链表尾部的节点，插入该节点并更新位置

type Node struct {
	key, value int
	pre, next  *Node
}

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*Node
	head, tail *Node
}

// InitNode 构建结点
func InitNode(key, value int) *Node {
	return &Node{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		cache:    map[int]*Node{},
		head:     InitNode(0, 0),
		tail:     InitNode(0, 0),
	}
	lru.head.next = lru.tail
	lru.tail.pre = lru.head
	return lru
}

// addToHead 结点添加到链表头部
func (this *LRUCache) addToHead(node *Node) {
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
}

// removeNode 删除结点
func (this *LRUCache) removeNode(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

// moveToHead 将链表中的结点移到头部
func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

// removeTail 删除尾结点
func (this *LRUCache) removeTail() *Node {
	node := this.tail.pre
	this.removeNode(node)
	return node
}

// Get 读取哈希表中的值、将改节点的位置移到链表头部（更新）
func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

// Put：如果缓存中已经存在该值，则更新值，更新节点在链表中的位置
//      如果缓存中该值不存在，则如果缓存中还有位置，则插入到头部；
//                          如果缓存中没有位置，则删去链表尾部的节点，插入该节点并更新位置
func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; ok {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	} else {
		node := InitNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			deleN := this.removeTail()
			// 别忘记删除map中的！！！！！并更新size大小
			delete(this.cache, deleN.key)
			this.size--
		}
	}
}

// 2022.11.19
