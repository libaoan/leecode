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

// todo 迭代法 未完成
func mergeTrees2(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	newRoot := &TreeNode{root1.Val + root2.Val, nil, nil}
	queue := []*TreeNode{newRoot}
	queue1 := []*TreeNode{root1}
	queue2 := []*TreeNode{root2}

	for len(queue1) > 0 || len(queue2) > 0 {
		node1 := queue1[0]
		queue1 = queue1[1:]
		node2 := queue2[0]
		queue2 = queue2[1:]
		node := queue[0]
		queue = queue[1:]

		if node1 == nil && node2 == nil {
			return newRoot
		}

		if node1 != nil && node2 != nil {
			node1l := node1.Left
			node1r := node1.Right
			node2l := node2.Left
			node2r := node2.Right
			newNodel := &TreeNode{node1l.Val + node2l.Val, nil, nil}
			newNoder := &TreeNode{node1r.Val + node2r.Val, nil, nil}
			node.Left = newNodel
			node.Right = newNoder
			queue = append(queue, newNodel, newNoder)
			queue1 = append(queue1, node1l, node1r)
			queue2 = append(queue2, node2l, node2r)
		}

		if node1 == nil {
			node2l := node2.Left
			node2r := node2.Right
			newNodel := &TreeNode{node2l.Val, nil, nil}
			newNoder := &TreeNode{node2r.Val, nil, nil}
			node.Left = newNodel
			node.Right = newNoder
			queue = append(queue, newNodel, newNoder)
			queue2 = append(queue2, node2l, node2r)
		} else {
			node1l := node1.Left
			node1r := node1.Right
			newNodel := &TreeNode{node1l.Val, nil, nil}
			newNoder := &TreeNode{node1r.Val, nil, nil}
			node.Left = newNodel
			node.Right = newNoder
			queue = append(queue, newNodel, newNoder)
			queue1 = append(queue1, node1l, node1r)
		}

	}

	return newRoot

}
