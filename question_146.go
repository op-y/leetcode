package leetcode

//import (
//	"fmt"
//)

type LRUNode struct {
	key  int
	val  int
	prev *LRUNode
	next *LRUNode
}

type LRUCache struct {
	capacity int
	n        int
	head     *LRUNode
	tail     *LRUNode
	index    map[int]*LRUNode
}

func Constructor146(capacity int) LRUCache {
	index := map[int]*LRUNode{}
	cache := LRUCache{
		capacity: capacity,
		index:    index,
	}
	return cache
}

func (this *LRUCache) Get(key int) int {
	if this.n <= 0 {
		return -1
	}
	if node, ok := this.index[key]; ok {
		// 找到后需要更新该元素位置
		// 元素就在队首情况无需调整
		if this.head == node {
			return node.val
		}
		// 其它情况需要调整位置
		node.prev.next = node.next // 前一个元素直接指向后一个元素
		if node.next != nil {
			// 当前元素不在尾部
			node.next.prev = node.prev
		} else {
			// 当前元素在尾部
			this.tail = node.prev
		}
		// 当前元素移动到队首
		node.next = this.head
		node.prev = nil
		this.head.prev = node
		this.head = node
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// 缓存中已经存在该元素
	if node, ok := this.index[key]; ok {
		node.val = value
		// 调整元素位置
		// 元素就在队首情况无需调整
		if this.head == node {
			return
		}
		// 其它情况需要调整位置
		node.prev.next = node.next // 前一个元素直接指向后一个元素
		if node.next != nil {
			// 当前元素不在尾部
			node.next.prev = node.prev
		} else {
			// 当前元素在尾部
			this.tail = node.prev
		}
		// 当前元素移动到队首
		node.next = this.head
		node.prev = nil
		this.head.prev = node
		this.head = node
		return
	}

	// 元素不在缓存中
	// 缓存已满 需要淘汰最旧的元素
	if this.n == this.capacity {
		last := this.tail
		// 只有一个元素的情况
		if this.n == 1 {
			this.head = nil
			this.tail = nil
		} else {
			last.prev.next = nil
			this.tail = last.prev
		}
		// 丢弃last元素
		last.prev = nil
		last.next = nil
		delete(this.index, last.key)
		this.n--
	}
	// 在队首加入新元素
	node := &LRUNode{
		key: key,
		val: value,
	}
	if this.n == 0 {
		this.head = node
		this.tail = node
	} else {
		node.next = this.head
		this.head.prev = node
		this.head = node
	}
	this.index[key] = node
	this.n++
	return
}

//func main() {
//	lRUCache := Constructor(2)
//	lRUCache.Put(2, 1)
//	lRUCache.Put(1, 1)
//	lRUCache.Put(2, 3)
//	lRUCache.Put(4, 1)
//	fmt.Printf("%d\n", lRUCache.Get(1))
//	fmt.Printf("%d\n", lRUCache.Get(2))
//}
