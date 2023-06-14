package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs T=41% S=50%
func isValidBST(root *TreeNode) bool {

	if root == nil {
		return true
	}
	return check(root, root.Left) && check(root, root.Right)

}

func check(root, cur *TreeNode) bool {

	if cur == nil {
		return true
	}

	if cur.Left != nil && cur.Left.Val >= cur.Val {
		return false
	}

	if cur.Right != nil && cur.Right.Val <= cur.Val {
		return false
	}

	return check(root, cur.Left) && check(root, cur.Right) && bst(root, cur)
}

func bst(root, cur *TreeNode) bool {

	if root == nil {
		return false
	}

	if root.Val < cur.Val {
		return bst(root.Right, cur)
	} else if root.Val > cur.Val {
		return bst(root.Left, cur)
	} else {
		if root != cur {
			return false
		} else {
			return true
		}
	}
}
