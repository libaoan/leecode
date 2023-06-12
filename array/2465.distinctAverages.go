package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 1, 4, 0, 3, 5}
	fmt.Println(distinctAverages(nums))
}

// O(n), 速度 100%， 内存 90%
func distinctAverages(nums []int) int {
	maps, i, j := make(map[int]struct{}, 0), 0, len(nums)-1
	ans := 0
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i < j {
		if _, ok := maps[nums[i]+nums[j]]; !ok {
			maps[nums[i]+nums[j]] = struct{}{}
			ans++
		}
		i++
		j--
	}
	return ans
}
