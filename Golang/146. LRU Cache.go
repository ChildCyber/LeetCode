package leetcode

// LRU 缓存
// https://leetcode-cn.com/problems/lru-cache/
type LRUNode struct {
	key, value int
	prev, next *LRUNode
}

type LRUCache struct {
	capacity   int
	cache      map[int]*LRUNode // 哈希表：提供 O(1) 的查找能力
	head, tail *LRUNode         // 双向链表：维护使用顺序，最近使用的在头部，最久未使用的在尾部
}

func Constructor146(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*LRUNode, capacity),
		head:     &LRUNode{key: -1, value: -1}, // 虚拟头节点
		tail:     &LRUNode{key: -1, value: -1}, // 虚拟尾节点
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.cache[key]; exists {
		// 移动到链表头部（表示最近使用）
		this.moveToHead(node)
		return node.value
	}
	// 不存在，返回-1
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// 键已存在，更新值并移动到头部
	if node, exists := this.cache[key]; exists {
		node.value = value
		this.moveToHead(node)
		return
	}

	// 键不存在，创建新节点
	newNode := &LRUNode{key: key, value: value}

	if len(this.cache) >= this.capacity {
		// 缓存已满，删除最久未使用的节点（尾部前一个）
		tailNode := this.tail.prev
		this.removeNode(tailNode)
		delete(this.cache, tailNode.key)
	}

	// 添加新节点到头部
	this.addToHead(newNode)
	this.cache[key] = newNode
}

// 将节点移动到链表头部
func (this *LRUCache) moveToHead(node *LRUNode) {
	// 先删除原位置，再插入新位置
	this.removeNode(node)
	this.addToHead(node)
}

// 从链表中删除节点
func (this *LRUCache) removeNode(node *LRUNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 添加节点到链表头部
func (this *LRUCache) addToHead(node *LRUNode) {
	// 当前节点操作
	node.prev = this.head
	node.next = this.head.next
	// 前后节点操作
	this.head.next.prev = node
	this.head.next = node
}
