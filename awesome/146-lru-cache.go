package awesome


//请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
//实现 LRUCache 类：
//LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
//int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
//void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
//函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

type LRUCache struct {
	values map[int]*Node
	capacity int  //容量
	len int   //长度

	// 双向链表
	head *Node
	tail *Node
}

type Node struct {
	key int
	val int
	pre *Node
	next *Node
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		values:   make(map[int]*Node, 1000),
		capacity: capacity,
		len:      0,
	}
}

// DeleteTail 删除尾巴
func (c *LRUCache) DeleteTail() {
	// 驱逐双向链表最后的一个
	if c.tail != nil {
		delete(c.values, c.tail.key)
	}
	// 切换尾巴
	c.tail = c.tail.pre
	if c.tail != nil {
		c.tail.next = nil
	}
}

// AddToHead 删除尾巴
func (c *LRUCache) AddToHead(node *Node) {
	// 拼接到开头
	tempHead := c.head

	if c.head == node {
		return
	}

	// 如果尾巴没有赋值，则直接赋值
	if c.tail == nil {
		c.tail = node
	}

	// 如果移动的是尾巴，则调整尾巴
	if node.next == nil && node.pre != nil{
		c.tail = node.pre
	}


	// 提取出来，上下拼接
	if node.pre != nil {
		node.pre.next = node.next
	}
	if node.next != nil {
		node.next.pre = node.pre
	}

	// 放置到开头
	c.head = node
	node.pre = nil
	node.next = tempHead

	if tempHead != nil {
		tempHead.pre = node
	}

}


func (c *LRUCache) Get(key int) int {
	node, ok := c.values[key]
	if !ok {
		return -1
	}

	// 拼接到开头
	c.AddToHead(node)

	return node.val
}


func (c *LRUCache) Put(key int, value int)  {
	node, ok := c.values[key]
	if ok {
		node.val = value  //值变更
		c.AddToHead(node)
	} else {

		newNode := &Node{
			key:  key,
			val:  value,
		}
		c.values[key] = newNode
		c.AddToHead(newNode)

		c.len = c.len + 1
	}

	// 大于容量
	if c.len > c.capacity {
		c.DeleteTail()
		c.len = c.len - 1
	}
}

//["LRUCache","put","put","get","get","put","get","get","get"]
//[[2],[2,1],[3,2],[3],[2],[4,3],[2],[3],[4]]
//输出：
//[null,null,null,2,1,null,-1,2,3]
//预期：
//[null,null,null,2,1,null,1,-1,3]


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */