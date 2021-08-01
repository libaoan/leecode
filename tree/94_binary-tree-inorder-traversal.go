package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// 递归算法
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	sres := inorderTraversal(root.Left)
	if sres != nil {
		res = append(res, sres...)
	}

	res = append(res, root.Val)

	sres = inorderTraversal(root.Right)
	if sres != nil {
		res = append(res, sres...)
	}

	return res
}

// 递归算法，优化空间
func inorderTraversal2(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}

// 迭代算法， 空间复杂度为2O(n)
func inorderTraversal3(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}

// todo Morris 中序遍历, 空间复杂度为O(1)
func inorderTraversal4(root *TreeNode) []int {

}
