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

// hashmap 速度 2N 25%, 内存 N 18%
func singleNumber2(nums []int) int {
	n := len(nums)
	maps := make(map[int]bool, n/2+1)
	for _, i := range nums {
		if _, ok := maps[i]; !ok {
			maps[i] = false
		} else {
			maps[i] = true
		}
	}

	for k, v := range maps {
		if !v {
			return k
		}
	}
	return -1
}

func main() {

}
