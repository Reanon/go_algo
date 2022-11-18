package lru_cache

const (
	hostbit = uint64(^uint(0)) == ^uint64(0)
	LENGTH  = 100
)

type lruNode struct {
	prev *lruNode
	next *lruNode
	// lru key
	key int
	// lru value
	value int
	hnext *lruNode // 拉链
}

type LRUCache struct {
	node []lruNode // hash list

	head *lruNode // lru head node
	tail *lruNode // lru tail node

	capacity int //
	used     int //
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		node:     make([]lruNode, LENGTH),
		head:     nil,
		tail:     nil,
		capacity: capacity,
		used:     0,
	}
}

func (c *LRUCache) Get(key int) int {
	if c.tail == nil {
		return -1
	}

	if tmp := c.searchNode(key); tmp != nil {
		c.moveToTail(tmp)
		return tmp.value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	// 1. 首次插入数据
	// 2. 插入数据不在 LRU 中
	// 3. 插入数据在 LRU 中
	// 4. 插入数据不在 LRU 中, 并且 LRU 已满
	if tmp := c.searchNode(key); tmp != nil {
		tmp.value = value
		c.moveToTail(tmp)
		return
	}
	c.addNode(key, value)

	if c.used > c.capacity {
		c.delNode()
	}
}

func (c *LRUCache) addNode(key int, value int) {
	newNode := &lruNode{
		key:   key,
		value: value,
	}

	tmp := &c.node[hash(key)]
	newNode.hnext = tmp.hnext
	tmp.hnext = newNode
	c.used++

	if c.tail == nil {
		c.tail, c.head = newNode, newNode
		return
	}
	c.tail.next = newNode
	newNode.prev = c.tail
	c.tail = newNode
}

func (c *LRUCache) delNode() {
	if c.head == nil {
		return
	}
	prev := &c.node[hash(c.head.key)]
	tmp := prev.hnext

	for tmp != nil && tmp.key != c.head.key {
		prev = tmp
		tmp = tmp.hnext
	}
	if tmp == nil {
		return
	}
	prev.hnext = tmp.hnext
	c.head = c.head.next
	c.head.prev = nil
	c.used--
}

func (c *LRUCache) searchNode(key int) *lruNode {
	if c.tail == nil {
		return nil
	}

	// 查找
	tmp := c.node[hash(key)].hnext
	for tmp != nil {
		if tmp.key == key {
			return tmp
		}
		tmp = tmp.hnext
	}
	return nil
}

func (c *LRUCache) moveToTail(node *lruNode) {
	if c.tail == node {
		return
	}
	if c.head == node {
		c.head = node.next
		c.head.prev = nil
	} else {
		node.next.prev = node.prev
		node.prev.next = node.next
	}

	node.next = nil
	c.tail.next = node
	node.prev = c.tail

	c.tail = node
}

func hash(key int) int {
	if hostbit {
		return (key ^ (key >> 32)) & (LENGTH - 1)
	}
	return (key ^ (key >> 16)) & (LENGTH - 1)
}
