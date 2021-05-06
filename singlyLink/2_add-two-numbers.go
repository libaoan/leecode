package main

func main() {
	l1 := &ListNode{}
	p1 := l1
	for _, v := range []int{2, 3, 4} {
		p1.Val = v
		p1.Next = &ListNode{}
		p1 = p1.Next
	}
	l2 := &ListNode{}
	p2 := l2
	for _, v := range []int{4, 5, 6} {
		p2.Val = v
		p2.Next = &ListNode{}
		p2 = p2.Next
	}

	l := addTwoNumbers(l1, l2)
	println(l)
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newL := &ListNode{}
	p := newL
	p1, p2 := l1, l2
	flag := 0
	for {
		if p1.Next != nil && p2.Next != nil {
			flag = p1.Val + p2.Val + flag
			if flag >= 10 {
				p.Val = flag % 10
				flag = flag - 10
			} else {
				p.Val = flag
				flag = 0
			}
			p.Next = &ListNode{}
			p = p.Next
			p1 = p1.Next
			p2 = p2.Next
		} else if p1.Next == nil && p2.Next == nil {
			return newL
		} else {
			break
		}
	}
	for {
		if p1.Next != nil {
			flag = p1.Val + flag
			if flag >= 10 {
				p.Val = flag % 10
				flag = flag - 10
			} else {
				p.Val = flag
				flag = 0
			}
			p.Next = &ListNode{}
			p = p.Next
			p1 = p1.Next
		} else if flag != 0 {
			p.Val = flag
			p.Next = &ListNode{}
			p = p.Next
			return newL
		} else {
			return newL
		}
		if p2.Next != nil {
			flag = p2.Val + flag
			if flag >= 10 {
				p.Val = flag % 10
				flag = flag - 10
			} else {
				p.Val = flag
				flag = 0
			}
			p.Next = &ListNode{}
			p = p.Next
			p2 = p2.Next
		} else if flag != 0 {
			p.Val = flag
			p.Next = &ListNode{}
			p = p.Next
			return newL
		} else {
			return newL
		}

		if flag != 0 {
			p.Val = flag
			p.Next = &ListNode{}
			p = p.Next
			return newL
		} else {
			return newL
		}

	}
}
