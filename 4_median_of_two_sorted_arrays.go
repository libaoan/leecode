package main

import "fmt"

func main(){
	fmt.Println(findMedianSortedArrays1([]int{}, []int{}))
	fmt.Println(findMedianSortedArrays1([]int{1,3,5}, []int{2,4,6}))
	fmt.Println(findMedianSortedArrays1([]int{1,3,5,7}, []int{2,3,6}))

}

// 合并两个数组后直接取中位数即可
// 时间复杂度 O(M+N) // not ok
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64{
	res := merge(nums1, nums2)
	n := len(res)
	if n == 0 {
		return -1
	}
	if n%2 == 0{
		return float64(res[n/2-1]+res[n/2]) / 2
	}

	return float64(res[n/2])
}

// Merge operate for two sorted nums
func merge(nums1, nums2 []int) []int{
	n1, n2 := len(nums2), len(nums2)
	i, j := 0, 0
	res := make([]int, 0, n1+n2)
	for i < n1 && j < n2{
		switch {
			case nums1[i] < nums2[j]:
				res = append(res, nums1[i])
				i++
			case nums1[i] > nums2[j]:
				res = append(res, nums2[j])
				j++
			default:
				res = append(res, nums1[i], nums2[j])
				i++
				j++
		}
	}

	if i < n1 {
		res = append(res, nums1[i:]...)
	}
	if j < n2 {
		res = append(res, nums2[j:]...)
	}
	return res
}

// 类二分查找分而治之的思路
// O(log(M+N)) // ok

