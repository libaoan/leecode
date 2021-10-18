package main

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	if root1 == nil && root2 == nil {
		return root1
	}
	if root1 != nil && root2 != nil {
		root1.Val += root2.Val
		root1.Left = mergeTrees(root1.Left, root2.Left)
		root1.Right = mergeTrees(root1.Right, root2.Right)
	}

	if root1 == nil {
		root1 = root2
		root1.Left = mergeTrees(root1.Left, nil)
		root1.Right = mergeTrees(root1.Right, nil)
	}

	if root2 == nil {
		root1.Left = mergeTrees(root1.Left, nil)
		root1.Right = mergeTrees(root1.Right, nil)
	}

	return root1
}
