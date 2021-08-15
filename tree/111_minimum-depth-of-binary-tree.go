package main

import "math"

// 递归法
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	mins := math.MaxInt32
	if root.Left != nil {
		mins = min(minDepth(root.Left)+1, mins)
	}
	if root.Right != nil {
		mins = min(minDepth(root.Right)+1, mins)
	}
	return mins

}

// 迭代法
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	depth := 1
	for len(queue) > 0 {

		cnt := len(queue)
		for cnt > 0 {
			node := queue[0]
			if cnt > 1 {
				queue = queue[1:]
			}
			if node.Left == nil && node.Right == nil {
				return depth
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			cnt--
		}
		depth++
	}

	return depth
}

func min(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}
