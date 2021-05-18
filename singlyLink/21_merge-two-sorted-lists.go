package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1, l2 ListNode
	p1, p2 := &l1, &l2
	num1 := []int{1, 2, 4}
	for i, v := range num1 {
		if i == 0 {
			l1 = ListNode{v, nil}
			p1 = &l1
		} else {
			p1.Next = &ListNode{v, nil}
			p1 = p1.Next
		}
	}
	num2 := []int{1, 3, 4}
	for i, v := range num2 {
		if i == 0 {
			l2 = ListNode{v, nil}
			p2 = &l2
		} else {
			p2.Next = &ListNode{v, nil}
			p2 = p2.Next
		}
	}

	//l := mergeTwoLists(&l1, &l2)
	l := mergeTwoLists2(&l1, &l2)
	for {
		if l != nil {
			fmt.Printf("%d ", l.Val)
			l = l.Next
		} else {
			fmt.Println()
			break
		}
	}

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2

	if p1 == nil {
		return p2
	}

	if p2 == nil {
		return p1
	}

	var head *ListNode
	if p1.Val <= p2.Val {
		head = p1
		p1 = p1.Next
	} else {
		head = p2
		p2 = p2.Next
	}
	p := head

	for {
		if p1 != nil && p2 != nil {
			if p1.Val > p2.Val {
				p.Next = p2
				p = p.Next
				p2 = p2.Next
			} else {
				p.Next = p1
				p = p.Next
				p1 = p1.Next
			}
		} else if p2 == nil {
			p.Next = p1
			return head
		} else {
			p.Next = p2
			return head
		}
	}
}

// 递归算法
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	if p1 == nil {
		return p2
	}

	if p2 == nil {
		return p1
	}

	var head *ListNode
	if p1.Val <= p2.Val {
		head = p1
		p1 = p1.Next
	} else {
		head = p2
		p2 = p2.Next
	}
	subList := mergeTwoLists2(p1, p2)
	head.Next = subList
	return head
}
