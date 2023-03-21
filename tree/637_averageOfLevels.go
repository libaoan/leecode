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

// 深度遍历， 速度 68% 内存 20%
func averageOfLevels2(root *TreeNode) []float64 {
	data := make([]struct {
		sum int
		cnt int
	}, 0)
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(data) > level {
			data[level].sum += node.Val
			data[level].cnt++
		} else {
			data = append(data, struct {
				sum int
				cnt int
			}{sum: node.Val, cnt: 1})
		}

		dfs(node.Right, level+1)
		dfs(node.Left, level+1)
	}
	dfs(root, 0)
	res := make([]float64, 0)
	for _, d := range data {
		res = append(res, float64(d.sum)/float64(d.cnt))
	}
	return res
}
