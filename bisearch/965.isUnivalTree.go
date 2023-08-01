package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// T(O) = 100% S(O) = 6%
func isUnivalTree(root *TreeNode) bool {
	li := []*TreeNode{}
	if root != nil {
		li = append(li, root)
	}
	for len(li) != 0 {
		node := li[0]
		if node.Val != root.Val {
			return false
		}
		if node.Left != nil {
			li = append(li, node.Left)
		}
		if node.Right != nil {
			li = append(li, node.Right)
		}
		if len(li) > 1 {
			li = li[1:]
		} else {
			break
		}
	}
	return true
}
