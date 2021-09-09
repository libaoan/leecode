package main

func sumOfLeftLeaves(root *TreeNode) int {
	sum := []int{0}
	sumOfLeftLeave(root, sum)
	return sum[0]
}

// 深度遍历，递归
func sumOfLeftLeave(root *TreeNode, sum []int) {
	if root == nil {
		return
	}

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum[0] += root.Left.Val
	}

	sumOfLeftLeave(root.Left, sum)
	sumOfLeftLeave(root.Right, sum)

	return

}

// 前序遍历 迭代法
func sumOfLeftLeaves2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			queue = []*TreeNode{}
		}
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			sum += node.Left.Val
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}

	}
	return sum
}

// 广度遍历， 迭代法
func sumOfLeftLeaves3(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		cnt := len(queue)

		for cnt > 0 {
			node := queue[0]
			if len(queue) > 1 {
				queue = queue[1:]
			} else {
				queue = []*TreeNode{}
			}
			if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
				sum += node.Left.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			cnt--

		}
	}
	return sum
}
