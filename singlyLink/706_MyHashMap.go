package main

type Data struct {
	value int
	next  *Data
}

type MyHashMap struct {
	data []Data
	cap  int
}

func Constructor() MyHashMap {
	return MyHashMap{}
}

func (this *MyHashMap) hash(key int) int {
	return key % this.cap
}

func (this *MyHashMap) Put(key int, value int) {
	if this.cap == 0 || this.cap <= this.hash(key) {
		cap_ := this.cap + key
		tmp := MyHashMap{
			data: make([]Data, cap_),
			cap:  cap_,
		}
		for k, v := range this.data {
			tmp.data[k] = v
		}
		this = &tmp
	}

	key = this.hash(key)
	da := Data{}
	if this.data[key] == da {
		da.value = value
		this.data[key] = da
	} else {
		cur := this.data[key]
		for cur.value != value && cur.next != nil {
			cur = *(cur.next)
		}
		if cur.value == value {
			return
		}
		da.value = value
		cur.next = &da
	}
}

func (this *MyHashMap) Get(key int) int {
	key = this.hash(key)
	if this.cap <= key {
		return -1
	}
}

func (this *MyHashMap) Remove(key int) {

}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
