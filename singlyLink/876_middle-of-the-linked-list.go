package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1 ListNode
	p1 := &l1
	num1 := []int{1, 2, 3, 4, 5, 6}
	for i, v := range num1 {
		if i == 0 {
			l1 = ListNode{v, nil}
			p1 = &l1
		} else {
			p1.Next = &ListNode{v, nil}
			p1 = p1.Next
		}
	}

	l := middleNode2(&l1)
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

// 快慢指针迭代法
func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for {
		if fast != nil && fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
		} else if fast != nil && fast.Next != nil {
			return slow.Next
		} else {
			return slow
		}
	}
}

// 快慢指针迭代法, 别人的代码
func middleNode2(head *ListNode) *ListNode {
	slow, fast := head, head
	for ; fast != nil && fast.Next != nil; fast = fast.Next.Next {
		slow = slow.Next
	}
	return slow
}
