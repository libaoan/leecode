package main

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	left := invertTree(root.Right)
	right := invertTree(root.Left)
	root.Right = left
	root.Left = right
	return root

}

// todo 迭代，广度优先
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			queue = []*TreeNode{}
		}

		left, right := node, node
		if node.Right != nil {
			left = node.Right
			queue = append(queue, node.Right)
		}

		if node.Left != nil {
			right = node.Left
			queue = append(queue, node.Left)
		}

		if left != node {
			node.Right = left
		} else {
			node.Right = nil
		}
		if right != node {
			node.Left = right
		} else {
			node.Left = nil
		}

	}
	return root
}
