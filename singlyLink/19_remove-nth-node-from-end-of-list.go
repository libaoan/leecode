package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1 ListNode
	p1 := &l1
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

	//l := removeNthFromEnd(&l1, 2)
	//l := removeNthFromEnd3(&l1, 2)
	l := removeNthFromEnd4(&l1, 2)
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

// 基本算法
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	p1, p2 := head, head
	midle, count := 0, 0

	for {
		if p1 != nil && p1.Next != nil && p1.Next.Next != nil {
			p1 = p1.Next.Next
			p2 = p2.Next
			count += 2
		} else if p1 != nil && p1.Next != nil {
			p1 = p1.Next.Next
			count += 2
			midle = count / 2
			break
		} else if p1 != nil {
			p1 = p1.Next
			count += 1
			midle = count/2 + 1
			break
		} else {
			break
		}
	}

	if count == 0 {
		return nil
	}
	if count < n {
		return head
	}
	index := count - n + 1

	if midle < index {
		for ; midle < index-1; midle++ {
			p2 = p2.Next
		}
		p2.Next = p2.Next.Next
	} else {
		if index == 1 {
			return head.Next
		} else {
			p1 = head
			for i := 1; i < index-1; i++ {
				p1 = p1.Next
			}
			p1.Next = p1.Next.Next
		}
	}
	return head

}

func getLegth(head *ListNode) int {
	p := head
	count := 0
	for p != nil {
		p = p.Next
		count++
	}
	return count
}

// 添加哑节点的二次遍历, 时间复杂度 n*n， 速度比我慢
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	length := getLegth(head)
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy

}

// 堆栈法，时间复杂度 n， 空间复杂度n， 空间复杂度比我差， 时间复杂度比我好
func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	stack := []*ListNode{}

	dumpy := &ListNode{0, head}
	for node := dumpy; node != nil; node = node.Next {
		stack = append(stack, node)
	}
	pre := stack[len(stack)-n-1]
	pre.Next = pre.Next.Next
	return dumpy.Next
}

// 快慢针法，空间和时间复杂度最优
func removeNthFromEnd4(head *ListNode, n int) *ListNode {
	dumpy := &ListNode{0, head}
	first, second := head, dumpy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dumpy.Next
}
