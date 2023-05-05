package main

import "fmt"

func main() {

	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}
	fmt.Println(intersect(nums1, nums2))
}

// O(n2), 速度 73%, 内存 88%
func intersect(nums1 []int, nums2 []int) []int {

	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	ans := []int{}

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				ans = append(ans, nums2[j])
				nums2[j] = -1
				break
			}
		}
	}

	return ans
}
