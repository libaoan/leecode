package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	tree := sortedArrayToBST(nums)
	fmt.Println(tree)
}

func sortedArrayToBST(nums []int) *TreeNode {

	if nums == nil {
		return nil
	}
	return Create(nums, 0, len(nums)-1)
}

func Create(nums []int, low, high int) *TreeNode {
	length := high - low
	if length < 0 {
		return nil
	}

	mid := low + length/2
	left := mid - 1
	right := mid + 1
	root := &TreeNode{nums[mid], nil, nil}
	if left >= low {
		root.Left = Create(nums, low, left)
	}

	if right <= high {
		root.Right = Create(nums, right, high)
	}

	return root
}
