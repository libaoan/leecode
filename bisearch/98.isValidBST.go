package main

func main() {

}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs todo 通过率 90%
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Left.Val >= root.Val {
		return false
	}
	if root.Right != nil && root.Right.Val <= root.Val {
		return false
	}

	return isValidBST(root.Left) && isValidBST(root.Right)
}
