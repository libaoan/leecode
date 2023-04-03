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

// 通过堆栈深度遍历， 速度 43% 内存 13%
func isSubtree2(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}

	if root == nil || subRoot == nil {
		return false
	}

	stack := []*TreeNode{root}
	targetNodes := []*TreeNode{}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		if node.Val == subRoot.Val {
			targetNodes = append(targetNodes, node)
		}
		for node.Left != nil {
			stack = append(stack, node.Left)
			node = node.Left
			if node.Val == subRoot.Val {
				targetNodes = append(targetNodes, node)
			}
		}
		stack = stack[:len(stack)-1]
		if node.Right != nil {
			stack = append(stack, node.Right)
			node = node.Right
			if node.Val == subRoot.Val {
				targetNodes = append(targetNodes, node)
			}
		} else {
			for len(stack) > 0 && node.Right == nil {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
				node = node.Right
				if node.Val == subRoot.Val {
					targetNodes = append(targetNodes, node)
				}
			}

		}

	}

	for _, node := range targetNodes {
		if deepEqual(node, subRoot) {
			return true
		}
	}
	return false
}
