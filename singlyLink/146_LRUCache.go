package main

import "fmt"

// todo: 通过率90%，本地测试无问题
type Node struct {
	Val  int
	Key  int
	Next *Node
}

type LRUCache struct {
	data *Node
	cap  int
}

func Constructor(capacity int) LRUCache {

	return LRUCache{cap: capacity}
}

func (this *LRUCache) Get(key int) int {

	if this.data == nil {
		return -1
	}

	pre, cur := this.data, this.data.Next
	if pre.Key == key {
		return pre.Val
	}
	for cur != nil {
		if cur.Key == key {
			pre.Next = cur.Next
			cur.Next = this.data
			this.data = cur
			return cur.Val
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	return -1
}

// 如果cap为0， 直接返回
// 遍历data，如果data为空，则添加元素
//   如果在首节点，直接更新元素
//   否则进行遍历，并检查是否超过容量
//     如果没有超过容量，并找到放到队首，并更新元素
//     如果没有超过容量，没有找到， 队首添加元素
//       如果超过容量，删除队尾的元素，队首添加新元素

func (this *LRUCache) Put(key int, value int) {

	if this.cap < 1 {
		return
	}
	if this.data == nil {
		this.data = &Node{Val: value, Key: key}
		return
	}

	pre, cur := this.data, this.data.Next
	if pre.Key == key {
		pre.Val = value
		return
	}

	cnt := 1
	for cur != nil {
		if cur.Key != key {
			cnt++
			pre = cur
			cur = cur.Next
		} else {
			pre.Next = cur.Next
			cur.Next = this.data
			this.data = cur
			cur.Val = value
			return
		}
	}

	if cnt < this.cap {
		this.data = &Node{Key: key, Val: value, Next: this.data}
	} else {
		pre, cur := this.data, this.data
		for cur.Next != nil {
			pre = cur
			cur = cur.Next
		}
		if pre.Next != nil {
			pre.Next = nil
			this.data = &Node{Key: key, Val: value, Next: this.data}
		} else {
			this.data = &Node{Key: key, Val: value}
		}
	}

}

func main() {

	lru := Constructor(100000)
	for i := 0; i < 100000; i++ {
		lru.Put(i, i)
		fmt.Println(lru.Get(i))
	}
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(2999))
}
