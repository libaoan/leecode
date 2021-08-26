package main

// 迭代法 todo: 很多问题
func postorderTraversal(root *TreeNode) []int {
	queue := []*TreeNode{}
	var res []int
	if root == nil {
		return res
	}
	for root != nil {
		queue = append(queue, root)
		root = root.Left
	}

	var prev *TreeNode

	for len(queue) > 0 {
		cnt := len(queue)
		node := queue[cnt-1]

		if len(queue) > 1 {
			queue = queue[0 : len(queue)-1]
		} else {
			queue = []*TreeNode{}
		}

		if node.Right == nil || node.Right == prev {
			res = append(res, node.Val)
			prev = node
			node = nil
		}
		if node != nil && node.Right != nil {
			queue = append(queue, node)
			queue = append(queue, node.Right)
			node = node.Right
			for node.Left != nil {
				queue = append(queue, node.Left)
				node = node.Left
			}
		}

	}

	return res

}

//  todo: 迭代法，别人的代码
func postorderTraversal2(root *TreeNode) (res []int) {
	stk := []*TreeNode{}
	var prev *TreeNode
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stk = append(stk, root)
			root = root.Right
		}
	}
	return
}
