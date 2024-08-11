package _146

type Node struct {
	prev  *Node
	next  *Node
	key   int
	value int
}

type LRUCache struct {
	head     *Node
	tail     *Node
	capacity int
	cache    map[int]*Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		head:     head,
		capacity: capacity,
		cache:    make(map[int]*Node),
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, found := this.cache[key]; found {
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, found := this.cache[key]; found {
		node.value = value
		this.moveToHead(node)
	} else {
		newNode := &Node{key: key, value: value}
		this.cache[key] = newNode
		this.addNode(newNode)

		if len(this.cache) > this.capacity {
			tail := this.popTail()
			delete(this.cache, tail.key)
		}
	}
}

func (this *LRUCache) addNode(node *Node) {
	node.prev = this.head
	node.next = this.head.next

	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *Node) {
	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev
}
func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addNode(node)
}

func (this *LRUCache) popTail() *Node {
	res := this.tail.prev
	this.removeNode(res)
	return res
}
