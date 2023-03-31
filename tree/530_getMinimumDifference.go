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
	queue := []*TreeNode{root}
	vals := []int{}
	for {
		node := queue[0]
		left, right := node.Left, node.Right
		switch {
		case left == nil && right == nil:
			break
		case left == nil:
			queue = append(queue, right)
			if node.Val == right.Val {
				return 0
			}
		case right == nil:
			queue = append(queue, left)
			if node.Val == left.Val {
				return 0
			}
		case left != nil && right != nil:
			queue = append(queue, left, right)
			if node.Val == right.Val || node.Val == left.Val {
				return 0
			}
		}

		vals = append(vals, node.Val)
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			break
		}
	}

	mins, ans := math.MaxInt, math.MaxInt
	for i := 0; i < len(vals); i++ {
		for j := i + 1; j < len(vals); j++ {
			if vals[i] == vals[j] {
				return 0
			}
			if vals[i] < vals[j] {
				mins = vals[j] - vals[i]
			} else {
				mins = vals[i] - vals[j]
			}
			if mins < ans {
				ans = mins
			}
		}
	}

	return ans
}

// todo: 优化深度遍历， 通过率50%
func getMinimumDifference2(root *TreeNode) int {

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
	mins := math.MaxInt

	for len(queue) > 0 {
		node := queue[len(queue)-1]
		// 遍历左子树
		for node.Left != nil {
			delta := abs(node.Val, node.Left.Val)
			if mins > delta {
				mins = delta
			}
			node = node.Left
			queue = append(queue, node)
		}
		// 左子树为空， node要出队列, 出队时，要遍历队列并计算差值
		queue = queue[:len(queue)-1]
		for _, n := range queue {
			delta := abs(node.Val, n.Val)
			if mins > delta {
				mins = delta
			}
		}

		// 遍历node的右子树
		//   如果右子树不为空，加入队列
		//   如果右子树为空，继续出队一个新的node，并检查它的右子树
		if node.Right != nil {
			queue = append(queue, node.Right)
			delta := abs(node.Val, node.Right.Val)
			if mins > delta {
				mins = delta
			}
		} else {
			// 如果node为叶子节点，从堆栈中弹出一个新的节点,并进行遍历求差
			for len(queue) > 0 && node.Right == nil {
				node = queue[len(queue)-1]
				queue = queue[:len(queue)-1]
				for _, n := range queue {
					delta := abs(node.Val, n.Val)
					if mins > delta {
						mins = delta
					}
				}
			}
			// 如果从堆栈弹出的节点右子树不为空，加入堆栈继续遍历
			if node.Right != nil {
				queue = append(queue, node.Right)
				delta := abs(node.Val, node.Right.Val)
				if mins > delta {
					mins = delta
				}
			}
		}
	}

	return mins
}
