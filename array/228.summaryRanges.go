package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	strs := summaryRanges(nums)
	for _, s := range strs {
		fmt.Println(s)
	}
}

// T(O) = 100% S(O) = 100%
func summaryRanges(nums []int) []string {
	strs, ln := []string{}, len(nums)
	for start := 0; start < ln; {
		end := ln - 1
		for start < ln && nums[start] <= nums[end] {
			if start == end {
				strs = append(strs, fmt.Sprintf("%d", nums[start]))
				start = end + 1
			} else if nums[end] == nums[start]+(end-start) {
				strs = append(strs, fmt.Sprintf("%d->%d", nums[start], nums[end]))
				start = end + 1
			} else {
				end--
			}
		}
	}

	return strs
}
