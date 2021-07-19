package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 0, 1}
	nums = []int{0, 1}
	// nums = []int{9,6,4,2,3,5,7,0,1}
	fmt.Println(missingNumber(nums))
	fmt.Println(missingNumber2(nums))
	fmt.Println(missingNumber3(nums))
}

// 哈希查找法
func missingNumber(nums []int) int {

	maps := make(map[int]bool)

	for i := range nums {
		maps[nums[i]] = true
	}

	target := 0
	for i := 0; i <= len(nums); i++ {
		if _, ok := maps[i]; !ok {
			target = i
			break
		}
	}
	return target
}

// 优化空间， 采用内置的排序算法排序，然后进行线性查找
func missingNumber2(nums []int) int {

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}
	return len(nums)
}

//  数学法，数列求和
func missingNumber3(nums []int) int {

	sumNs := 0

	ln := len(nums)
	sumN := ln * (ln + 1) / 2

	for i := range nums {
		sumNs += nums[i]
	}

	return sumN - sumNs
}
