package main

// 迭代
func searchBST(root *TreeNode, val int) *TreeNode {

	for root != nil && root.Val != val {
		if root.Val < val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}

// 递归
func searchBST2(root *TreeNode, val int) *TreeNode {

	switch {
	case root == nil || root.Val == val:
		return root
	case root.Val < val:
		return searchBST2(root.Right, val)
	case root.Val > val:
		return searchBST2(root.Left, val)
	}
	return nil
}
