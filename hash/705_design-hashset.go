package main

import (
	"container/list"
	"fmt"
)

func main() {
	hs := Constructor()
	hs.Add(1)                   // set = [1]
	hs.Add(2)                   // set = [1, 2]
	fmt.Println(hs.Contains(1)) // 返回 True
	fmt.Println(hs.Contains(3)) // 返回 False ，（未找到）
	hs.Add(2)                   // set = [1, 2]
	fmt.Println(hs.Contains(2)) // 返回 True
	hs.Remove(2)                // set = [1]
	fmt.Println(hs.Contains(2)) // 返回 False ，（已移除）

	fmt.Println("*****")
	hs2 := ConstructorMyHashSetWithOpenAddr()
	hs2.Add(1)                   // set = [1]
	hs2.Add(2)                   // set = [1, 2]
	fmt.Println(hs2.Contains(1)) // 返回 True
	fmt.Println(hs2.Contains(3)) // 返回 False ，（未找到）
	hs2.Add(2)                   // set = [1, 2]
	fmt.Println(hs2.Contains(2)) // 返回 True
	hs2.Remove(2)                // set = [1]
	fmt.Println(hs2.Contains(2)) // 返回 False ，（已移除）
	fmt.Println(hs2.Contains(0))
}

const base = 769

// 链表法
type MyHashSetUseList struct {
	data []list.List
}

func Constructor() MyHashSetUseList {
	return MyHashSetUseList{make([]list.List, base)}
}

func (s *MyHashSetUseList) hash(key int) int {
	return key % base
}

func (s *MyHashSetUseList) Add(key int) {
	if !s.Contains(key) {
		h := s.hash(key)
		s.data[h].PushBack(key)
	}
}

func (s *MyHashSetUseList) Remove(key int) {
	h := s.hash(key)
	for e := s.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			s.data[h].Remove(e)
		}
	}
}

func (s *MyHashSetUseList) Contains(key int) bool {
	h := s.hash(key)
	for e := s.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			return true
		}
	}
	return false
}

const empty = -1
const deleted = -2

// 开放地址法
type MyHashSetWithOpenAddr struct {
	caps, count int
	data        []int
}

func ConstructorMyHashSetWithOpenAddr() MyHashSetWithOpenAddr {
	ns := MyHashSetWithOpenAddr{base, 0, make([]int, base)}
	for i := 0; i < ns.caps; i++ {
		ns.data[i] = empty
	}
	return ns
}

func (s *MyHashSetWithOpenAddr) hash(key int) int {
	return key % s.caps
}

func (s *MyHashSetWithOpenAddr) Add(key int) {
	if !s.Contains(key) {
		if float32(s.count)/float32(s.caps) > 0.75 {
			ns := &MyHashSetWithOpenAddr{2 * s.caps, 0, make([]int, 2*s.caps)}
			for i := 0; i < ns.caps; i++ {
				ns.data[i] = empty
			}
			for i := 0; i < s.caps; i++ {
				if s.data[i] != empty && s.data[i] != deleted {
					ns.Add(s.data[i])
				}
			}
			s = ns
		}
		h := s.hash(key)
		for i := 0; i < s.caps; i++ {
			if s.data[(h+i)%s.caps] == empty {
				s.data[(h+i)%s.caps] = key
				s.count++
				break
			}
		}
	}
}

func (s *MyHashSetWithOpenAddr) Remove(key int) {
	if s.Contains(key) {
		if float32(s.count)/float32(s.caps) < 0.001 {
			ns := &MyHashSetWithOpenAddr{s.caps / 2, 0, make([]int, s.caps/2)}
			for i := 0; i < ns.caps; i++ {
				ns.data[i] = empty
			}
			for i := 0; i < s.caps; i++ {
				if s.data[i] != empty && s.data[i] != deleted {
					ns.Add(s.data[i])
				}
			}
			s = ns
		}
		h := s.hash(key)
		for i := 0; i < s.caps; i++ {
			if s.data[(h+i)%s.caps] == key {
				s.data[(h+i)%s.caps] = deleted
				s.count--
				break
			}
		}
	}

}

func (s *MyHashSetWithOpenAddr) Contains(key int) bool {
	h := s.hash(key)
	for i := 0; i < s.caps; i++ {
		if s.data[(h+i)%s.caps] == key {
			return true
		}
	}
	return false
}

// todo 二次hash法
type MyHashSetWithDoubleHash struct {
	data []int
}
