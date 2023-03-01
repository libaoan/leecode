package main

import (
	"fmt"
	"sort"
)

// 贪心 时间复杂度n
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	res := 0

	for i := 0; i < len(nums); i += 2 {
		if nums[i] < nums[i+1] {
			res += nums[i]
		} else {
			res += nums[i+1]
		}
	}

	return res
}

func main() {
	nums := []int{6, 2, 6, 5, 1, 2}
	fmt.Println(arrayPairSum(nums))
}
