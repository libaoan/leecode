package main

import "fmt"

func main() {
	l1 := &ListNode{}
	p1 := l1
	num1 := []int{2, 4, 3}
	lenth1 := len(num1)
	for i, v := range num1 {
		p1.Val = v
		if i != lenth1-1 {
			p1.Next = &ListNode{}
			p1 = p1.Next
		}
	}
	l2 := &ListNode{}
	p2 := l2
	num2 := []int{5, 6, 7}
	lenth2 := len(num2)
	for i, v := range num2 {
		p2.Val = v
		if i != lenth2-1 {
			p2.Next = &ListNode{}
			p2 = p2.Next
		}
	}

	l := addTwoNumbers(l1, l2)
	fmt.Println("%s", *l)
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
		if p1 != nil && p2 != nil {
			flag = p1.Val + p2.Val + flag
			if flag >= 10 {
				p.Val = flag - 10
				flag = flag / 10
			} else {
				p.Val = flag
				flag = 0
			}
			p.Next = &ListNode{}
			p = p.Next
			p1 = p1.Next
			p2 = p2.Next
		} else if p1 == nil && p2 == nil {
			if flag != 0 {
				p.Val = flag
				return newL
			} else {
				return newL
			}
		} else {
			break
		}
	}
	for {
		if p1 != nil {
			flag = p1.Val + flag
			if flag >= 10 {
				p.Val = flag - 10
				flag = flag / 10
			} else {
				p.Val = flag
				flag = 0
			}
			p.Next = &ListNode{}
			p = p.Next
			p1 = p1.Next
		} else {
			return newL
		}
		if p2 != nil {
			flag = p2.Val + flag
			if flag >= 10 {
				p.Val = flag - 10
				flag = flag / 10
			} else {
				p.Val = flag
				flag = 0
			}
			p.Next = &ListNode{}
			p = p.Next
			p2 = p2.Next
		} else {
			return newL
		}

	}
}
