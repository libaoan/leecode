package main

import "math"

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 空间 n 时间n2  速度3% 内存 3%
// todo 待优化
func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	vals := []int{}
	for {
		node := queue[0]
		left, right := node.Left, node.Right
		switch {
		case left == nil && right == nil:
			break
		case left == nil:
			queue = append(queue, right)
			if node.Val == right.Val {
				return 0
			}
		case right == nil:
			queue = append(queue, left)
			if node.Val == left.Val {
				return 0
			}
		case left != nil && right != nil:
			queue = append(queue, left, right)
			if node.Val == right.Val || node.Val == left.Val {
				return 0
			}
		}

		vals = append(vals, node.Val)
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			break
		}
	}

	mins, ans := math.MaxInt, math.MaxInt
	for i := 0; i < len(vals); i++ {
		for j := i + 1; j < len(vals); j++ {
			if vals[i] == vals[j] {
				return 0
			}
			if vals[i] < vals[j] {
				mins = vals[j] - vals[i]
			} else {
				mins = vals[i] - vals[j]
			}
			if mins < ans {
				ans = mins
			}
		}
	}

	return ans
}
