package main

// 两次遍历
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	getPath := func(root *TreeNode, target int) []*TreeNode {
		res := []*TreeNode{}
		for root != nil {
			if root.Val == target {
				res = append(res, root)
				return res
			}
			if root.Val < target {
				res = append(res, root)
				root = root.Right
			} else {
				res = append(res, root)
				root = root.Left
			}
		}
		return nil
	}

	node := root
	pPath := getPath(node, p.Val)
	node = root
	qPath := getPath(node, q.Val)

	if pPath == nil || qPath == nil {
		return nil
	}

	if len(pPath) < len(qPath) {
		pPath, qPath = qPath, pPath
	}
	targetNode := root
	for i := 0; i < len(qPath); i++ {
		if pPath[i].Val == qPath[i].Val {
			targetNode = qPath[i]
		} else {
			return targetNode
		}
	}

	return qPath[len(qPath)-1]
}

// todo 优化：一次遍历
func lowestCommonAncestor2(root, p, q *TreeNode) (ancestor *TreeNode) {
	ancestor = root
	for {
		if p.Val < ancestor.Val && q.Val < ancestor.Val {
			ancestor = ancestor.Left
		} else if p.Val > ancestor.Val && q.Val > ancestor.Val {
			ancestor = ancestor.Right
		} else {
			return
		}
	}
}
