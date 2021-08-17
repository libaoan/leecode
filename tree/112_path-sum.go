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

// 迭代法
func hasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	root.Val = targetSum - root.Val
	queue := []*TreeNode{root}

	for len(queue) > 0 {

		cnt := len(queue)
		for cnt > 0 {
			node := queue[0]
			leftVal := node.Val
			if node.Left == nil && node.Right == nil && node.Val == 0 {
				return true
			}

			if node.Left != nil {
				node.Left.Val = leftVal - node.Left.Val
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				node.Right.Val = leftVal - node.Right.Val
				queue = append(queue, node.Right)
			}
			if len(queue) > 1 {
				queue = queue[1:]
			} else {
				return false
			}
			cnt--
		}

	}

	return false
}
