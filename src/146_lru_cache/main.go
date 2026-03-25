package main

type Node struct {
	value int
	next  *Node
	prev  *Node
	key   int
}
type LRUCache struct {
	capacity  int
	dummyHead *Node
	dummyTail *Node
	counter   int
	nodeMap   map[int]*Node
}

func Constructor(capacity int) LRUCache {
	dummyHead := &Node{}
	dummyTail := &Node{}
	dummyHead.next = dummyTail
	dummyTail.prev = dummyHead
	return LRUCache{
		capacity:  capacity,
		counter:   0,
		nodeMap:   make(map[int]*Node),
		dummyHead: dummyHead,
		dummyTail: dummyTail,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.nodeMap[key]
	if !ok {
		return -1
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = this.dummyTail.prev
	this.dummyTail.prev.next = node
	node.next = this.dummyTail
	this.dummyTail.prev = node
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.nodeMap[key]
	if !ok {
		node = &Node{value: value, key: key}
		this.nodeMap[key] = node
		this.counter++
	} else {
		node.value = value
		node.next.prev = node.prev
		node.prev.next = node.next
	}
	node.prev = this.dummyTail.prev
	this.dummyTail.prev.next = node
	node.next = this.dummyTail
	this.dummyTail.prev = node
	if this.counter > this.capacity {
		delete(this.nodeMap, this.dummyHead.next.key)
		this.dummyHead.next.next.prev = this.dummyHead
		this.dummyHead.next = this.dummyHead.next.next
		this.counter--
	}

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
