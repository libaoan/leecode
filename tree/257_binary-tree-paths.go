package main

import (
	"k8s.io/client-go/tools/cache"
	"strconv"
	"strings"
)

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	res := []string{}
	path := []*TreeNode{root}
	if root.Left != nil {
		getPath(root.Left, path)
	}

}

func getPath(root *TreeNode, path []*TreeNode) []*TreeNode {
	path = append(path, root)
	if root.Left == nil && root.Right == nil {
		return path
	}
	if root.Left != nil {
		return getPath(root.Left, path)
	}
	if root.Right != nil {
		return getPath(root.Right, path)
	}
}

// 深度迭代遍历 todo 待优化
func binaryTreePaths2(root *TreeNode) []string {
	res := []string{}

	queue := []*TreeNode{}

	for root != nil {
		queue = append(queue, root)
		root = root.Left
	}
	prev := root
	for len(queue) > 0 {
		node := queue[len(queue)-1]

		if node.Right == nil && node.Left == nil {
			path := []string{}
			for _, p := range queue {
				path = append(path, strconv.Itoa(p.Val))
			}
			res = append(res, strings.Join(path, "->"))
		}

		if len(queue) > 1 {
			queue = queue[:len(queue)-1]
		} else {
			queue = []*TreeNode{}
		}

		if node.Right != nil && node.Right != prev {
			queue = append(queue, node)
			prev = node.Right
			queue = append(queue, prev)
			node = prev
			for node.Left != nil {
				queue = append(queue, node.Left)
				node = node.Left
			}
		} else {
			prev = node
		}
	}
	return res
}
