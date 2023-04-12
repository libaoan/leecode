package main

import "sort"

// 通过， 速度 12% 内存62%
func singleNumber(nums []int) int {
	sort.Ints(nums)
	i := 0
	for ; i < len(nums)-1; i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[i]
}

func main() {

}
