package main

import "fmt"

const Cap = 673

type Data struct {
	value int
	key   int
	next  *Data
}

// todo 通过率90%
type MyHashMap struct {
	data [Cap]*Data
}

func Constructor() MyHashMap {
	return MyHashMap{
		data: [Cap]*Data{},
	}
}

func (this *MyHashMap) hash(key int) int {
	return key % Cap
}

func (this *MyHashMap) Put(key int, value int) {
	k := this.hash(key)
	pre, cur := this.data[k], this.data[k]
	for cur != nil && cur.key != key {
		pre = cur
		cur = cur.next
	}
	// empty
	if pre == nil {
		this.data[k] = &Data{key: key, value: value}
		return
	}
	// found and replace
	if cur != nil {
		cur.value = value
		return
	}
	// not found
	pre.next = &Data{key: key, value: value}
}

func (this *MyHashMap) Get(key int) int {
	k := this.hash(key)
	cur := this.data[k]
	for cur != nil && cur.key != key {
		cur = cur.next
	}
	if cur == nil {
		return -1
	}
	return cur.value
}

func (this *MyHashMap) Remove(key int) {
	k := this.hash(key)
	pre, cur := this.data[k], this.data[k]
	for cur != nil && cur.key != key {
		pre = cur
		cur = cur.next
	}
	if cur == nil {
		return
	}
	if pre.next == nil {
		this.data[k] = nil
	} else {
		pre.next = cur.next
		cur = nil
	}
}

func (this *MyHashMap) printAll() {
	for i := range this.data {
		cur := this.data[i]
		// fmt.Printf("xxx %T, %v, %t\r\n", cur, cur, cur == nil)
		if cur == nil {
			continue
		}

		for ; cur != nil; cur = cur.next {
			fmt.Printf("[%d, %d]", cur.key, cur.value)
		}
		fmt.Println()
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

func main() {
	obj := Constructor()
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.printAll()
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(2))
	obj.Put(2, 1)
	obj.printAll()
	//fmt.Println(obj.data[0].next)
	fmt.Println(obj.Get(2))
}
