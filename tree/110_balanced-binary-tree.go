package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	nums := []string{"null", "1", "2", "2", "3", "3", "null", "null", "4", "4"}
	nums = []string{"2", "1", "4", "null", "null", "3", "6", "null", "null", "5"}
	nums = []string{"1", "2"}

	root := CreateTree(nums)

	fmt.Println(isBalanced(root))
}

// 构造二叉树
func CreateTree(nums []string) *TreeNode {
	if len(nums) == 0 || nums[0] == "null" {
		return nil
	}
	queue := []*TreeNode{}
	num, _ := strconv.Atoi(nums[0])
	root := &TreeNode{num, nil, nil}
	queue = append(queue, root)
	for index := 0; index < len(nums) && len(queue) > 0; index += 2 {
		node := queue[0]
		if len(queue) > 0 {
			queue = queue[1:]
		}

		if index+1 < len(nums) && nums[index+1] != "null" {
			num, _ := strconv.Atoi(nums[index+1])
			newn := &TreeNode{num, nil, nil}
			node.Left = newn
			queue = append(queue, newn)
		}
		if index+2 < len(nums) && nums[index+2] != "null" {
			num, _ := strconv.Atoi(nums[index+2])
			newn := &TreeNode{num, nil, nil}
			node.Right = newn
			queue = append(queue, newn)
		}
	}
	return root
}

func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}

	leftD := getDeep2(root.Left) + 1
	rightD := getDeep2(root.Right) + 1

	if leftD-rightD > 1 || rightD-leftD > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

// 递归算法
func getDeep(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftD := getDeep(root.Left) + 1
	rightD := getDeep(root.Right) + 1

	if leftD >= rightD {
		return leftD
	} else {
		return rightD
	}
}

// 迭代算法
func getDeep2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	depth := 0

	for len(queue) > 0 {

		subL := len(queue)
		for subL > 0 {
			node := queue[0]
			if len(queue) > 0 {
				queue = queue[1:]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			subL--
		}
		depth++
	}
	return depth
}
