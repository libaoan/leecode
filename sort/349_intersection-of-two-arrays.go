package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	nums := intersection2(nums1, nums2)
	fmt.Println(nums)
}

// 基础算法，hash表， 复杂度（n+m）,空间复杂度 2n
func intersection(nums1 []int, nums2 []int) []int {
	maps := make(map[int]bool)

	for _, n := range nums1 {
		maps[n] = false
	}

	for _, n := range nums2 {
		if _, ok := maps[n]; ok {
			maps[n] = true
		}
	}

	nums := []int{}
	for k, v := range maps {
		if v {
			nums = append(nums, k)
		}
	}
	return nums

}

// 有序列表 + 双指针算法， 空间复杂度 n
func intersection2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	nums := []int{}

	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] == nums2[j] {
			if len(nums) == 0 || nums1[i] != nums[len(nums)-1] {
				nums = append(nums, nums1[i])
			}
			i++
			j++
		} else {
			j++
		}
	}

	return nums
}
