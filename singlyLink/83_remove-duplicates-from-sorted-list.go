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

	//l := deleteDuplicates(&l1)
	l := deleteDuplicates2(&l1)
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

// 基础算法(时间复杂度为n，空间复杂度为n)， 没有利用有序性， 浪费较多额外空间(N)
func deleteDuplicates(head *ListNode) *ListNode {
	maps := make(map[int]bool)
	p, q := head, head
	if p == nil {
		return head
	} else {
		q = p
		maps[q.Val] = true
		p = p.Next
	}
	for ; p != nil; p = p.Next {
		ok := maps[p.Val]
		if ok {
			q.Next = p.Next
		} else {
			maps[p.Val] = true
			q = q.Next
		}
	}
	return head
}

// 优化，利用有序性
func deleteDuplicates2(head *ListNode) *ListNode {

	p := head
	for ; p != nil; p = p.Next {
		cur := p
		for p.Next != nil && cur.Val == p.Next.Val {
			p = p.Next
		}
		cur.Next = p.Next
	}
	return head

}
