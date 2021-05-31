package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1 ListNode
	p1 := &l1
	num1 := []int{1, 2, 2, 1}
	for i, v := range num1 {
		if i == 0 {
			l1 = ListNode{v, nil}
			p1 = &l1
		} else {
			p1.Next = &ListNode{v, nil}
			p1 = p1.Next
		}
	}

	l := isPalindrome(&l1)

	fmt.Printf("isPalindrome:%s ", l)

}

func isPalindrome(head *ListNode) bool {

	if head == nil {
		return true
	}

	p, slow, fast := head, head, head

	for {
		if fast != nil && fast.Next != nil && fast.Next.Next != nil && fast.Next.Next.Next != nil {
			fast = fast.Next.Next
			p = slow.Next
			slow.Next = slow.Next.Next
			p.Next = slow
		} else if fast != nil && fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
		} else {
			break
		}
	}

	slow = slow.Next
	p = head
	for {
		if slow != nil && p.Val == slow.Val {
			slow = slow.Next
			p = p.Next
		} else if slow != nil {
			return false
		} else {
			return true
		}
	}

}
