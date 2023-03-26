package main

import (
	"math"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 速度 100%， 内存38%
func getDecimalValue(head *ListNode) int {
	ans := []byte{}
	for p := head; p != nil; p = p.Next {
		ans = append(ans, byte(p.Val))
	}

	bits := len(ans) - 1
	res := 0
	for i, b := range ans {
		if b == 0 {
			continue
		}
		res += int(math.Pow(2, float64(bits-i)))
	}
	return res

}
