package main

import "fmt"

func main() {

	nums, target := []int{-1, 0, 3, 5, 9, 12}, 9
	nums, target = []int{-1, 0, 3, 5, 9, 12}, 2
	nums, target = []int{2, 5}, 5
	fmt.Println(search(nums, target))
}

// O(logn) 速度 97%, 内存 60%
func search(nums []int, target int) int {
	lo, hi, mid := 0, len(nums)-1, -1
	for lo <= hi {
		if nums[lo] == target {
			return lo
		}
		if nums[hi] == target {
			return hi
		}

		tt := lo + (hi-lo)/2
		if tt == mid {
			break
		}
		mid = tt
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			lo = mid
		} else {
			hi = mid
		}
		//fmt.Println(lo, hi, mid)
	}
	return -1
}
