package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// O(n2) 速度 14%， 内存 67%
func mergeKLists(lists []*ListNode) *ListNode {

	n := len(lists)
	if n == 0 {
		return nil
	}

	var p, head *ListNode
	for {
		k := 0
		for ; k < n; k++ {
			if lists[k] != nil {
				break
			}
		}
		if k == n {
			return head
		}
		for j := k + 1; j < n; j++ {
			if lists[j] != nil && lists[k].Val > lists[j].Val {
				k = j
			}
		}

		if head == nil {
			head = lists[k]
			lists[k] = lists[k].Next
			p = head
		} else {
			p.Next = lists[k]
			p = p.Next
			lists[k] = lists[k].Next
		}
	}

}
