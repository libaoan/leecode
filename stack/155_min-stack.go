package main

import (
	"fmt"
	"math"
)

func main() {
	obj := Constructor()
	obj.Push(2147483646)
	obj.Print()
	obj.Push(2147483646)
	obj.Print()
	obj.Push(2147483647)
	obj.Print()
	fmt.Println("getTop:", obj.Top())
	obj.Pop()
	fmt.Print("after pop:")
	obj.Print()
	fmt.Println("getMin:", obj.GetMin())
	obj.Pop()
	fmt.Print("after pop:")
	obj.Print()
	obj.Push(2147483647)
	obj.Print()
	fmt.Println("getTop:", obj.Top())
	fmt.Println("getMin:", obj.GetMin())
	obj.Push(-2147483648)
	obj.Print()
	fmt.Println("getTop:", obj.Top())
	fmt.Println("getMin:", obj.GetMin())
	obj.Pop()
	fmt.Print("after pop:")
	obj.Print()
	fmt.Println("getMin:", obj.GetMin())

}

// 算法复杂度O(n)
type MinStack struct {
	stack    []int
	n        int
	count    int
	minIndex int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{nil, 0, 0, -1}
}

func (this *MinStack) Print() {
	for i := 0; i < this.count; i++ {
		fmt.Print(this.stack[i], " ")
	}
	fmt.Println()
}

func (this *MinStack) Push(val int) {
	if this.count == this.n {
		this.n++
		this.count++
		this.stack = append(this.stack, val)
	} else {
		this.count++
		this.stack[this.count-1] = val
	}

	if this.minIndex == -1 || this.stack[this.minIndex] > val {
		this.minIndex = this.count - 1
	}
}

func (this *MinStack) Pop() {
	if this.count != 0 {
		if this.minIndex == this.count-1 {
			this.count--
			this.minIndex = this.count - 1
			for i := 0; i < this.count; i++ {
				if this.stack[this.minIndex] > this.stack[i] {
					this.minIndex = i
				}
			}
		} else {
			this.count--
		}
	}

}

func (this *MinStack) Top() int {
	if this.count != 0 {
		return this.stack[this.count-1]
	} else {
		return 0
	}
}

func (this *MinStack) GetMin() int {
	if this.count != 0 {
		return this.stack[this.minIndex]
	} else {
		return 0
	}
}

// 别人的代码，最小栈， 算法复杂度为O(1)
type MinStack2 struct {
	stack    []int
	minStack []int
}

func Constructor2() MinStack2 {
	return MinStack2{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack2) Push(x int) {
	this.stack = append(this.stack, x)
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(x, top))
}

func (this *MinStack2) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack2) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack2) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
