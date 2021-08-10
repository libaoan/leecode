package main

func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}

	leftD := getDeep(root.Left)
	rightD := getDeep(root.Right)

	if leftD-rightD > 1 || rightD-leftD > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

// 递归算法
func getDeep(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftD := getDeep(root.Left) + 1
	rightD := getDeep(root.Right) + 1

	if leftD >= rightD {
		return leftD
	} else {
		return rightD
	}
}

// todo 迭代算法
func getDeep2(root *TreeNode) int {

}
