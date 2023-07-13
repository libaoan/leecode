package main

// Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

// T(O) = 60% S(O) = 100%
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	p := head
	for p.Next != nil && p.Next.Val != val {
		p = p.Next
	}
	if p.Next == nil {
		return head
	} else {
		p.Next = p.Next.Next
		return head
	}

}
