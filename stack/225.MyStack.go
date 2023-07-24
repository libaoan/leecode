package main

// T(O) = 3% S(O) = 100%
type MyStack struct {
	stack []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyStack) Pop() int {
	ln := len(this.stack)
	if ln != 0 {
		p := this.stack[ln-1]
		this.stack = this.stack[:ln-1]
		return p
	} else {
		return 0
	}
}

func (this *MyStack) Top() int {
	ln := len(this.stack)
	if ln != 0 {
		p := this.stack[ln-1]
		return p
	} else {
		return 0
	}
}

func (this *MyStack) Empty() bool {
	ln := len(this.stack)
	if ln != 0 {
		return false
	} else {
		return true
	}
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
