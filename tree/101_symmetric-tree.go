package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	li := []int{1, 2, 2, 3, 4, 4, 3}
	root := CreateTree(li, 0)
	fmt.Println(isSymmetric2(root))
}

func CreateTree(li []int, index int) *TreeNode {
	root := &TreeNode{li[index], nil, nil}

	if len(li) > 2*index+1 {
		root.Left = CreateTree(li, 2*index+1)
	}
	if len(li) > 2*index+2 {
		root.Right = CreateTree(li, 2*index+2)
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
	queue = append(queue, p, q)
	for len(queue) > 0 {
		p := queue[0]
		q := queue[1]
		if len(queue) > 2 {
			queue = queue[2:]
		} else {
			queue = nil
		}

		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		queue = append(queue, p.Right, q.Left, p.Left, q.Right)
	}
	return true
}
