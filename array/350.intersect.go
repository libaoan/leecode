package main

import "fmt"

func main() {

	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}
	fmt.Println(intersect(nums1, nums2))
}

// O(n3), 通过率 50% todo： 审题错误，不考虑顺序
func intersect(nums1 []int, nums2 []int) []int {

	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	sets := [][]int{}

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); {
			if nums1[i] == nums2[j] {
				k := 1
				// fmt.Println(i, j, k)
				for i+k < len(nums1) && j+k < len(nums2) && nums1[i+k] == nums2[j+k] {
					k++
				}
				sets = append(sets, nums1[i:i+k])
				j = j + k
			} else {
				j++
			}
		}
	}

	ans := []int{}
	for _, set := range sets {
		if len(ans) < len(set) {
			ans = set
		}
	}
	return ans
}
