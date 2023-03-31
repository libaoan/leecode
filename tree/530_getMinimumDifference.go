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

// todo: 深度遍历， 通过率50%
func getMinimumDifference2(root *TreeNode) int {

	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	mins := math.MaxInt
	node := root
	// 左树
	for node != nil {
		for node.Left != nil {
			if mins > (node.Val - node.Left.Val) {
				mins = node.Val - node.Left.Val
			}
			node = node.Left
			queue = append(queue, node)
		}

		if len(queue) == 0 {
			break
		}

		if node.Right != nil {
			if mins > node.Right.Val-node.Val {
				mins = node.Right.Val - node.Val
			}
			if mins > root.Val-node.Val {
				mins = root.Val - node.Val
			}
			if len(queue) < 1 {
				break
			}
			queue = queue[:len(queue)-1]
			node = node.Right
			queue = append(queue, node)
		} else {
			if mins > root.Val-node.Val {
				mins = root.Val - node.Val
			}
			if len(queue) < 2 {
				break
			}
			queue = queue[:len(queue)-1]
			node = queue[len(queue)-1]
			if mins > root.Val-node.Val {
				mins = root.Val - node.Val
			}
			if node.Right != nil {
				if mins > node.Right.Val-node.Val {
					mins = node.Right.Val - node.Val
				}
				if mins > root.Val-node.Val {
					mins = root.Val - node.Val
				}
				if len(queue) < 2 {
					break
				}
				queue = queue[:len(queue)-1]
				queue = append(queue, node.Right)
			}
			node = node.Right
		}
	}

	// 右树
	node = root.Right
	queue = []*TreeNode{node}
	if node == nil {
		return mins
	}
	if mins > node.Val-root.Val {
		mins = node.Val - root.Val
	}
	for node != nil {
		for node.Left != nil {
			if mins > node.Val-node.Left.Val {
				mins = node.Val - node.Left.Val
			}
			node = node.Left
			queue = append(queue, node)
		}

		if len(queue) == 0 {
			break
		}

		if node.Right != nil {
			if mins > node.Right.Val-node.Val {
				mins = node.Right.Val - node.Val
			}
			if mins > node.Val-root.Val {
				mins = node.Val - root.Val
			}
			queue = queue[:len(queue)-1]
			node = node.Right
			queue = append(queue, node)
		} else {
			if mins > node.Val-root.Val {
				mins = node.Val - root.Val
			}
			if len(queue) < 2 {
				break
			}
			queue = queue[:len(queue)-1]
			node = queue[len(queue)-1]
			if node.Right != nil {
				if mins > node.Right.Val-node.Val {
					mins = node.Right.Val - node.Val
				}
				if mins > node.Val-root.Val {
					mins = node.Val - root.Val
				}
				if len(queue) < 2 {
					break
				}
				queue = queue[:len(queue)-1]
				queue = append(queue, node.Right)
			}
			node = node.Right
		}
	}
	return mins
}
