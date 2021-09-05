package main

import (
	"strconv"
	"strings"
)

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

// 广度迭代遍历 todo:用例通过率 98%
func binaryTreePaths4(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	res := [][]string{}
	queue := []*TreeNode{root}
	paths := [][]string{[]string{strconv.Itoa(root.Val)}}

	for len(queue) > 0 {
		cnt := len(queue)
		for cnt > 0 {
			node := queue[0]
			path := paths[0]
			if len(queue) > 1 {
				queue = queue[1:]
				paths = paths[1:]
			} else {
				queue = []*TreeNode{}
				paths = [][]string{}
			}
			if node.Left == nil && node.Right == nil {
				res = append(res, path)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
				// todo: append(path,...) 弱引用，和下面一个条件可能冲突 [6,1,null,null,3,2,5,null,null,4] failed
				tmp_path := append(path, strconv.Itoa(node.Left.Val))
				paths = append(paths, tmp_path)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				tmp_path := append(path, strconv.Itoa(node.Right.Val))
				paths = append(paths, tmp_path)
			}
			cnt--
		}
	}

	ress := []string{}
	for _, r := range res {
		ress = append(ress, strings.Join(r, "->"))
	}
	return ress
}
