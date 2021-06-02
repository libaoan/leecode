package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Print()
	obj.Push(2)
	obj.Print()
	fmt.Println("Peek:", obj.Peek())
	fmt.Println("Pop:", obj.Pop())
	obj.Print()
	fmt.Println("isEmpty:", obj.Empty())
	fmt.Println("Pop:", obj.Pop())
	obj.Print()
	fmt.Println("isEmpty:", obj.Empty())
}

type MyQueue struct {
	inStack, outStack []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.outStack) == 0 {
		this.inToOut()
	}
	item := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return item
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		this.inToOut()
	}
	return this.outStack[len(this.outStack)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}

func (this *MyQueue) inToOut() {
	for len(this.inStack) > 0 {
		this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
		this.inStack = this.inStack[:len(this.inStack)-1]
	}
}

func (this *MyQueue) Print() {
	this.inToOut()
	for i := 0; i < len(this.outStack); i++ {
		fmt.Print(this.outStack[i], " ")
	}
	fmt.Println()
}
