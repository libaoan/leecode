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

	l := reverseList2(&l1)
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

// 迭代法， 时间复杂度O(n)， 空间复杂度O(1)
func reverseList(head *ListNode) *ListNode {

	var second *ListNode
	p := head
	for p != nil {
		first := p
		p = p.Next
		first.Next = second
		second = first
	}

	return second
}

// 递归法 时间复杂度O(n) 空间复杂度O(n)
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
