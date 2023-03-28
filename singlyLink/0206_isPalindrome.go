package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// todo: 快慢指针
func isPalindrome(head *ListNode) bool {

	p, q := head, head
	var r *ListNode
	for q != nil && q.Next != nil {
		r = p
		p = p.Next
		q = q.Next.Next
		r.Next = head
		head = r
	}

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
