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

// 进阶算法 todo:参考

func findMode1(root *TreeNode) []int {
	res := []int{}
	var base, count, maxCount int

	update := func(x int) {
		if x == base {
			count++
		} else {
			base, count = x, 1
		}
		if count == maxCount {
			res = append(res, base)
		} else if count > maxCount {
			maxCount = count
			res = []int{base}
		}
	}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		update(node.Val)
		dfs(node.Right)
	}
	dfs(root)

	return res
}

// 迭代法，前序遍历  todo: 用例通过率90%
func findMode2(root *TreeNode) []int {
	var base, count, maxcount int = 0, 0, 0
	var res []int

	update := func(x int) {
		if x == base {
			count++
		} else {
			base, count = x, 1
		}
		if count == maxcount {
			res = append(res, base)
		} else if count > maxcount {
			maxcount = count
			res = []int{base}
		}
	}

	queue := []*TreeNode{}
	if root != nil {

		update(root.Val)
		queue = append(queue, root)
	}

	var pop *TreeNode
	for len(queue) > 0 {
		length := len(queue)
		node := queue[length-1]

		if node.Left == nil || node.Left != nil && node.Left == pop {

			length := len(queue)
			if length > 1 {
				queue = queue[0 : length-1]
			} else {
				queue = []*TreeNode{}
			}
		}

		for node.Left != nil && node.Left != pop {
			update(node.Left.Val)
			node = node.Left
		}

		if node.Right != nil {
			update(node.Right.Val)
			queue = append(queue, node.Right)
		} else {
			pop = node
			length := len(queue)
			if length > 1 {
				queue = queue[0 : length-1]
			} else {
				queue = []*TreeNode{}
			}
		}
	}

	return res

}
