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

// 删除尾巴
func (this *LRUCache) DeleteTail() {
	// 驱逐双向链表最后的一个
	if this.tail != nil {
		delete(this.values, this.tail.key)
	}

	// 切换尾巴
	this.tail = this.tail.pre
	if this.tail != nil {
		this.tail.next = nil
	}
}

// 删除尾巴
func (this *LRUCache) AddToHead(node *Node) {
	// 拼接到开头
	tempHead := this.head

	if this.head == node {
		return
	}

	// 如果尾巴没有赋值，则直接赋值
	if this.tail == nil {
		this.tail = node
	}

	// 如果移动的是尾巴，则调整尾巴
	if node.next == nil && node.pre != nil{
		this.tail = node.pre
	}


	// 提取出来，上下拼接
	if node.pre != nil {
		node.pre.next = node.next
	}
	if node.next != nil {
		node.next.pre = node.pre
	}

	// 放置到开头
	this.head = node
	node.pre = nil
	node.next = tempHead

	if tempHead != nil {
		tempHead.pre = node
	}

}


func (this *LRUCache) Get(key int) int {
	node, ok := this.values[key]
	if !ok {
		return -1
	}

	// 拼接到开头
	this.AddToHead(node)

	return node.val
}


func (this *LRUCache) Put(key int, value int)  {
	node, ok := this.values[key]
	if ok {
		node.val = value  //值变更
		this.AddToHead(node)
	} else {

		newNode := &Node{
			key:  key,
			val:  value,
		}
		this.values[key] = newNode
		this.AddToHead(newNode)

		this.len = this.len + 1
	}

	// 大于容量
	if this.len > this.capacity {
		this.DeleteTail()
		this.len = this.len - 1
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