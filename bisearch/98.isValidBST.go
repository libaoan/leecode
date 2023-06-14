package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 回朔 T=41% S=50%
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

// dfs T=100% S=6%
func isValidBST2(root *TreeNode) bool {

	if root == nil {
		return true
	}

	stack := []*TreeNode{}
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}

	li := []int{}

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		if node.Right != nil {
			tmp := node.Right
			stack = append(stack, tmp)
			for tmp.Left != nil {
				tmp = tmp.Left
				stack = append(stack, tmp)
			}
		}

		li = append(li, node.Val)
	}

	for i := 0; i < len(li)-1; i++ {
		if li[i] >= li[i+1] {
			return false
		}
	}
	return true

}
