package main

// 迭代，字典序
func findMode(root *TreeNode) []int {
	maps := make(map[int]int, 0)
	visitAndCount(root, maps)
	res := []int{}
	maxc := 0

	for k, v := range maps {
		switch {
		case maxc == v:
			res = append(res, k)
		case maxc < v:
			maxc = v
			res = []int{k}
		}
	}
	return res
}

func visitAndCount(root *TreeNode, maps map[int]int) {
	if root == nil {
		return
	}
	if _, ok := maps[root.Val]; ok {
		maps[root.Val]++
	} else {
		maps[root.Val] = 1
	}
	visitAndCount(root.Right, maps)
	visitAndCount(root.Left, maps)
}

// 进阶算法
