package main

func sumOfLeftLeaves(root *TreeNode) int {
	sum := []int{0}
	sumOfLeftLeave(root, sum)
	return sum[0]
}

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
