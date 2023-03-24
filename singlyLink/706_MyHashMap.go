package main

import "fmt"

type Data struct {
	value int
	key   int
	next  *Data
}

type MyHashMap struct {
	data []*Data
	cap  int
}

func Constructor() MyHashMap {
	return MyHashMap{}
}

func (this *MyHashMap) hash(key int) int {
	if this.cap == 0 {
		return key
	}
	return key % this.cap
}

func (this *MyHashMap) Put(key int, value int) {

	if this.cap == 0 || this.cap <= this.hash(key) {
		cap_ := this.cap + this.hash(key)
		tmp := MyHashMap{
			data: make([]*Data, cap_),
			cap:  cap_,
		}
		for _, v := range this.data {
			for v != nil {
				tmp.Put(v.key, v.value)
				v = v.next
			}
		}
		*this = tmp
	}

	k := this.hash(key)
	pre, cur := this.data[k], this.data[k]
	for cur != nil && cur.key != key {
		pre = cur
		cur = cur.next
	}
	if pre == nil {
		da := &Data{key: key, value: value}
		this.data[k] = da
		return
	}
	if cur != nil && cur.key == key {
		cur.value = value
		return
	}
	da := &Data{key: key, value: value}
	da.value = value
	da.key = key
	pre.next = da
}

func (this *MyHashMap) Get(key int) int {
	k := this.hash(key)
	if this.cap <= k {
		return -1
	}
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
	if this.cap <= k {
		return
	}
	pre, cur := this.data[k], this.data[k]
	for cur != nil && cur.key != key {
		pre = cur
		cur = cur.next
	}
	if cur == nil {
		return
	}
	pre.next = cur.next
	cur.next = nil
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
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(2))
	obj.Put(2, 1)
	//fmt.Println(obj.data[0].next)
	fmt.Println(obj.Get(2))
}
