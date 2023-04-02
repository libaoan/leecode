package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归算法，速度 43%, 空间 7%
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}
	nodes := search(root, subRoot.Val)
	for _, node := range nodes {
		if deepEqual(node, subRoot) {
			return true
		}
	}

	return false
}

func search(root *TreeNode, val int) []*TreeNode {
	if root == nil {
		return nil
	}
	res := []*TreeNode{}
	if root.Val == val {
		res = append(res, root)
	}
	node := search(root.Left, val)
	if node != nil {
		res = append(res, node...)
	}
	node = search(root.Right, val)
	if node != nil {
		res = append(res, node...)
	}
	return res
}

func deepEqual(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	if node1.Val != node2.Val {
		return false
	}
	if deepEqual(node1.Left, node2.Left) && deepEqual(node1.Right, node2.Right) {
		return true
	}
	return false
}
