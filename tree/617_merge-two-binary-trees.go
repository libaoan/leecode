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

// todo 迭代法
func mergeTrees2(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	p1, p2 := root1, root2
	p := p1
	if p1 == nil {
		p = p2
	}
	leftq, rightq := []*TreeNode{}, []*TreeNode{}
	if p1.Left != nil {
		leftq = append(leftq, p1.Left)
	}

	if p2.Left != nil {
		leftq = append(leftq, p2.Left)
	}

	if p1.Right != nil {
		rightq = append(rightq, p1.Right)
	}

	if p2.Right != nil {
		rightq = append(rightq, p2.Right)
	}

	for len(leftq) != 0 || len(rightq) != 0 {

		if p1 != nil && p2 != nil {
			p.Val = p1.Val + p2.Val

		}

	}

	for root1 != nil || root2 != nil {
		if root1 != nil && root2 != nil {
			root1.Val += root2.Val
			p1l = root1.Left
			p1r = root1.Right

			p2l = root2.Left
			p2r = root2.Right
		}
	}

}
