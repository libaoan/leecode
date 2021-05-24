package main

import (
	"flag"
	"fmt"
)

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

	flag := hasCycle2(&l1)
	fmt.Println(flag)

}

func hasCycle(head *ListNode) bool {

	p := head
	var maps = make(map[*ListNode]bool)
	for {
		if p != nil {
			ok := maps[p]
			if ok {
				return true
			} else {
				maps[p] = true
				p = p.Next
			}
		} else {
			return false
		}
	}

}

// 快慢指针法
func hasCycle2(head *ListNode) bool {
	p, q := head, head
	if p == nil {
		return false
	} else {
		q = p.Next
	}

	for p != nil && q != nil && q.Next != nil && p != q {
		q = q.Next.Next
		p = p.Next
	}
	if p != nil && p == q {
		return true
	} else {
		return false
	}
}

// 别人的代码
func hasCycle3(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
