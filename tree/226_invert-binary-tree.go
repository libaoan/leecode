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

// 迭代，广度优先
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			queue = []*TreeNode{}
		}

		left := node.Left
		right := node.Right

		node.Right = left
		node.Left = right

		if left != nil {
			queue = append(queue, left)
		}

		if right != nil {
			queue = append(queue, right)
		}

	}
	return root

}
