package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 0, -1, 0, -2, 2, -2}
	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	n := removeElement(nums, 2)
	fmt.Println("After:", n, nums)
}

// 遍历
func removeElement(nums []int, val int) int {
	left := 0
	for _, v := range nums { // v 即 nums[right]
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}
