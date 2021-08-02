package main

// 递归法
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && nil == q
	}

	if !isSameTree(p.Right, q.Right) {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	if !isSameTree(p.Left, q.Left) {
		return false
	}

	return true
}
