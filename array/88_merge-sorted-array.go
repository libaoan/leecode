package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge2(nums1, 6, nums2, 3)
	fmt.Println(nums1)
}

// TODO 本地调试通过
func merge2(nums1 []int, m int, nums2 []int, n int) {

	newNums := make([]int, m)
	i, j, k := 0, 0, 0
	for {
		if i >= m-n || j < n && nums1[i] >= nums2[j] {
			newNums[k] = nums2[j]
			j++
			k++
		} else {
			newNums[k] = nums1[i]
			i++
			k++
		}

		if k == m {
			break
		}
	}
	for i, v := range newNums {
		nums1[i] = v
	}
}
