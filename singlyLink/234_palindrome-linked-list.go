package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1 ListNode
	p1 := &l1
	num1 := []int{1, 3, 2, 4, 3, 2, 1}
	//num1 := []int{1, 2, 2, 1}
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

// todo: 不是最优算法
func isPalindrome(head *ListNode) bool {

	if head == nil {
		return true
	}
	head = &ListNode{0, head}
	p, slow, fast := head, head.Next, head.Next
	for {
		if fast != nil && fast.Next != nil && fast.Next.Next != nil && fast.Next.Next.Next != nil {
			fast = fast.Next.Next
			p = slow.Next
			slow.Next = slow.Next.Next
			p.Next = head.Next
			head.Next = p
		} else if fast != nil && fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
			break
		} else {
			break
		}
	}

	slow = slow.Next
	p = head.Next
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
