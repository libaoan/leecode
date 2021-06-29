package main

import "sort"

func main() {
	nums := []int{4, 2, 5, 9, 7, 4, 8}
	println(maxProductDifference(nums))
}

func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	lt := len(nums)
	return nums[lt-2]*nums[lt-1] - nums[0]*nums[1]
}
