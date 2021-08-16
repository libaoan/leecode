package main

// 递归
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	if hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val) {
		return true
	} else {
		return false
	}

}

// todo: 迭代法
func hasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		sum := node.Val
		if len(queue) > 1 {
			queue = queue[1:]
		}
		for node.Left != nil {
			queue = append(queue, node.Left)
			node = node.Left
			sum += node.Val
		}
		for node.Right != nil {
			queue = append(queue, node.Right)
			node = node.Right
			sum += node.Val
		}
		if sum == targetSum {
			return true
		}

	}

	return false
}
