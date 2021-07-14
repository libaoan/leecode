package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 2, 3}
	nums = []int{2, 2, 1, 1, 1, 2, 2}
	nums = []int{10, 9, 9, 9, 10}

	fmt.Println(majorityElement(nums))
	fmt.Println(majorityElement2(nums))
}

// 哈希表算法，空间复杂度n, 时间复杂度n
func majorityElement(nums []int) int {
	maps := make(map[int]int)

	for _, v := range nums {
		if cnt, ok := maps[v]; ok {
			maps[v] = cnt + 1
		} else {
			maps[v] = 1
		}
	}

	mid := len(nums) >> 1
	for k, v := range maps {
		if v > mid {
			return k
		}
	}
	// never executed
	return 0
}

// 空间复杂度1， 时间复杂度n
func majorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
