package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 6, 2, 3}
	fmt.Println(largestPerimeter(nums))
}

// 基础算法，已经最优
func largestPerimeter(nums []int) int {
	sort.Ints(nums)

	length := len(nums)
	if length < 3 {
		return 0
	}

	for i := length - 1; i > 1; i-- {
		if nums[i-2]+nums[i-1] > nums[i] {
			return nums[i] + nums[i-2] + nums[i-1]
		}
	}
	return 0
}
