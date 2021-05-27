package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1 ListNode
	p1 := &l1
	num1 := []int{1, 1, 2, 3, 3}
	for i, v := range num1 {
		if i == 0 {
			l1 = ListNode{v, nil}
			p1 = &l1
		} else {
			p1.Next = &ListNode{v, nil}
			p1 = p1.Next
		}
	}

	l := removeElements(&l1, 3)
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

// 已经用到最优算法啦
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	dump := &ListNode{0, head}

	pre, after := dump, dump.Next
	for ; after != nil; after = after.Next {
		if after.Val == val {
			pre.Next = after.Next
		} else {
			pre = pre.Next
		}
	}
	return dump.Next
}
