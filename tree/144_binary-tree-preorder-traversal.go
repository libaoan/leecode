package main

// 递归法
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{root.Val}
	if root.Left != nil {
		res = append(res, preorderTraversal(root.Left)...)
	}

	if root.Right != nil {
		res = append(res, preorderTraversal(root.Right)...)
	}

	return res
}

// 迭代法
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := []int{root.Val}
	queue := []*TreeNode{root}

	for root != nil && root.Left != nil {
		res = append(res, root.Left.Val)
		queue = append(queue, root.Left)
		root = root.Left
	}

	for len(queue) > 0 {
		node := queue[len(queue)-1]
		if len(queue) > 1 {
			queue = queue[:len(queue)-1]
		} else {
			queue = []*TreeNode{}
		}

		if node.Right != nil {
			res = append(res, node.Right.Val)
			queue = append(queue, node.Right)
			node = node.Right
			for node != nil && node.Left != nil {
				res = append(res, node.Left.Val)
				queue = append(queue, node.Left)
				node = node.Left
			}
		}
	}
	return res

}
