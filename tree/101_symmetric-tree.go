package main

import "fmt"

func main() {
	li := []int{1, 2, 2, 3, 4, 4, 3}
	root := CreateTree(li)
	fmt.Println(isSymmetric2(root))
}

// todo 构造二叉树问题
func CreateTree(li []int) *TreeNode {
	var root *TreeNode

	if len(li) != 0 {
		root = &TreeNode{li[0], nil, nil}
		root.Left = CreateTree(li[1:])
		root.Right = CreateTree(li[2:])
	}
	return root
}

// 递归
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

// 迭代

func isSymmetric2(root *TreeNode) bool {
	p, q := root, root
	queue := []*TreeNode{}
	queue = append(queue, p)
	queue = append(queue, q)
	for len(queue) >= 0 {
		p := queue[0]
		q := queue[1]
		queue := queue[2:]
		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		queue = append(queue, p.Right)
		queue = append(queue, q.Left)
		queue = append(queue, p.Left)
		queue = append(queue, q.Right)
	}
	return true
}
