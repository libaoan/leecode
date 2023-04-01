package main

import "math"

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 空间 n 时间n2  速度3% 内存 3%
// todo 待优化
func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}

	abs := func(i, j int) int {
		if i > j {
			return i - j
		}
		return j - i
	}

	queue := []*TreeNode{root}
	poped := []*TreeNode{root}
	mins := math.MaxInt

	for len(queue) > 0 {
		node := queue[len(queue)-1]
		// 遍历左子树
		for node.Left != nil {
			node = node.Left
			queue = append(queue, node)
		}
		// 左子树为空， node要出队列, 出队时，要遍历队列并计算差值
		queue = queue[:len(queue)-1]
		for _, n := range poped {
			if node == n {
				continue
			}
			delta := abs(node.Val, n.Val)
			if mins > delta {
				mins = delta
			}
		}
		poped = append(poped, node)

		// 遍历node的右子树
		//   如果右子树不为空，加入队列
		//   如果右子树为空，继续出队一个新的node，并检查它的右子树
		if node.Right != nil {
			queue = append(queue, node.Right)
		} else {
			// 如果node为叶子节点，从堆栈中弹出一个新的节点,并进行遍历求差
			for len(queue) > 0 && node.Right == nil {
				node = queue[len(queue)-1]
				queue = queue[:len(queue)-1]
				for _, n := range poped {
					if node == n {
						continue
					}
					delta := abs(node.Val, n.Val)
					if mins > delta {
						mins = delta
					}
				}
				poped = append(poped, node)
			}
			// 如果从堆栈弹出的节点右子树不为空，加入堆栈继续遍历
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if len(queue) == 0 {
				for _, n := range poped {
					if node == n {
						continue
					}
					delta := abs(node.Val, n.Val)
					if mins > delta {
						mins = delta
					}
				}
				poped = append(poped, node)
			}
		}
	}

	return mins
}
