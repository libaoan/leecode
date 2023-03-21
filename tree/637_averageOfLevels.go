package main

////Definition for a binary tree node.
//type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
//}

// 宽度遍历， 速度 68%， 内存 73%
func averageOfLevels(root *TreeNode) []float64 {
	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue, root)
	}
	res := make([]float64, 0)
	for len(queue) > 0 {
		var sum float64
		newQueue := make([]*TreeNode, 0)
		for _, node := range queue {
			sum += float64(node.Val)
			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
		}
		res = append(res, sum/float64(len(queue)))
		queue = newQueue
	}
	return res
}
