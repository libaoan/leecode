package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针 时间复杂度n 空间复杂度1 速度97%， 内存98%
func isPalindrome(head *ListNode) bool {

	if head == nil {
		return true
	}

	p, q := head, head
	var r *ListNode
	head = nil
	for q != nil {
		r = p
		p = p.Next
		if q.Next == nil {
			break
		}
		q = q.Next.Next

		r.Next = head
		head = r
	}
	r = head
	for p != nil && r != nil {
		if r.Val != p.Val {
			return false
		}
		p = p.Next
		r = r.Next
	}
	return true

}

// 速度 60%， 内存50% 空间复杂度n 时间复杂度n
func isPalindrome1(head *ListNode) bool {
	slice := []int{}
	for p := head; p != nil; p = p.Next {
		slice = append(slice, p.Val)
	}
	n := len(slice)

	for i := 0; i < n/2; i++ {
		if slice[i] != slice[n-i-1] {
			return false
		}
	}
	return true
}
