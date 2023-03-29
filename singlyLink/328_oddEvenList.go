package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 空间复杂度 1 ，时间复杂度 n 速度 80%， 空间100%
func oddEvenList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	odd, eve := head, head.Next
	p, q := odd, eve
	for q != nil {
		p.Next = q.Next
		if p.Next != nil {
			p = p.Next
			q.Next = p.Next
			q = q.Next
		} else {
			q.Next = nil
			q = q.Next
		}
	}
	p.Next = eve
	return odd

}
