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
	nums = []string{"null", "2", "1", "4", "null", "null", "3", "6", "null", "null", "5"}

	// todo: 构造失败
	root := CreateTree(nums)

	fmt.Println(isBalanced(root))
}

func CreateTree(nums []string) *TreeNode {
	if len(nums) < 1 {
		return nil
	}
	return Create(nums, 1)
}

func Create(nums []string, index int) *TreeNode {

	if nums[index] == "null" {
		return nil
	}

	num, _ := strconv.Atoi(nums[index])
	root := &TreeNode{num, nil, nil}
	leftI, rightI := 2*index, 2*index+1
	if leftI < len(nums) && nums[leftI] != "null" {
		root.Left = Create(nums, leftI)
	}

	if rightI < len(nums) && nums[rightI] != "null" {
		root.Right = Create(nums, rightI)
	}

	return root
}

func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}

	leftD := getDeep2(root.Left)
	rightD := getDeep2(root.Right)

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

// todo 迭代算法

func getDeep2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left, right := 0, 0

	for root.Left != nil {
		root = root.Left
		left++
	}

	for root.Right != nil {
		root = root.Right
		right++
	}
	if left >= right {
		return left + 1
	} else {
		return right + 1
	}
}
