package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归算法
func findSecondMinimumValue(root *TreeNode) int {
	target := -1
	if root.Right != nil && root.Val != root.Right.Val {
		target = root.Right.Val
	}

	if root.Left != nil && root.Val != root.Left.Val {
		target = root.Left.Val
	}

	target = getTarget(root, root.Val, target)

	return target
}

func getTarget(node *TreeNode, min, t int) int {
	if node == nil {
		return t
	}
	if node.Val != min {
		if t == -1 || t > node.Val {
			t = node.Val
		}
	}
	t1 := getTarget(node.Right, min, t)
	t2 := getTarget(node.Left, min, t)

	if t1 == -1 {
		return t2
	}
	if t2 == -1 {
		return t1
	}

	if t1 < t2 {
		return t1
	}
	return t2
}
