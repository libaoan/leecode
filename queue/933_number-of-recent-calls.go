package main

import "fmt"

func main() {
	s := []int{1, 100, 3001, 3002}
	this := Constructor()
	for _, n := range s {
		fmt.Printf("%d  ", this.Ping(n))
	}
}

type RecentCounter struct {
	queue      []int
	head, tail int
}

func Constructor() RecentCounter {
	return RecentCounter{make([]int, 3001), 0, 0}
}

func (this *RecentCounter) Ping(t int) int {

	// is this.queue full? deque if yes
	if this.tail-this.head >= 3001 {
		for t > 3000 && this.queue[this.head] < t-3000 {
			this.head++
		}
		if this.head == 0 {
			return -1
		}
	}

	// is this.queque ended? move if yes
	if this.tail >= 3001 {
		head := this.head
		tail := this.tail
		this.head = 0
		this.tail = 0
		for i := head; i < tail; i++ {
			this.queue[this.tail] = this.queue[i]
			this.tail++
		}
	}
	this.queue[this.tail] = t
	this.tail++
	for t > 3000 && this.queue[this.head] < t-3000 {
		this.head++
	}
	return this.tail - this.head
}
