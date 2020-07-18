package leetcode

type LRUCache struct {
	head, tail *LRUNode
	capacity   int
	len        int
	lruMap     map[int]*LRUNode
}

type LRUNode struct {
	key, value int
	next, prev *LRUNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		head:     &LRUNode{0, 0, nil, nil},
		tail:     &LRUNode{0, 0, nil, nil},
		capacity: capacity,
		lruMap:   make(map[int]*LRUNode, capacity),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.lruMap[key]
	if ok {
		this.moveToHead(v)
		return v.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	v, ok := this.lruMap[key]
	if ok {
		v.value = value
		this.moveToHead(v)
	} else {
		new := &LRUNode{
			key:   key,
			value: value,
		}
		this.lruMap[key] = new
		this.addToHead(new)
		this.len++

		if this.capacity < this.len {
			removed := this.removeTail()
			delete(this.lruMap, removed.key)
			this.len--
		}
	}

}

func (this *LRUCache) addToHead(node *LRUNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *LRUNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *LRUNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *LRUNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
