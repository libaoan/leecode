package main

import "fmt"

func main(){
	fmt.Println(findMedianSortedArrays1([]int{}, []int{}))
	fmt.Println(findMedianSortedArrays1([]int{1,3,5}, []int{2,4,6}))
	fmt.Println(findMedianSortedArrays1([]int{1,3,5,7}, []int{2,3,6}))
	fmt.Println("******************")
	fmt.Println(findMedianSortedArrays2([]int{}, []int{}))
	fmt.Println(findMedianSortedArrays2([]int{1,3,5}, []int{2,4,6}))
	fmt.Println(findMedianSortedArrays2([]int{1,3,5,7}, []int{2,3,6}))

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

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64{
	n1, n2 := len(nums1), len(nums2)
	if n1 + n2 == 0 {
		return -1
	}
	if (n1+n2)%2 ==0 {
		l := findKth(nums1, nums2, (n1+n2)/2)
		r := findKth(nums1, nums2, (n1+n2)/2+1)
		return float64(l+r)/2
	}
	return float64(findKth(nums1, nums2, (n1+n2)/2+1))
}

// 在两个有序数组中查找第 k 大的数
// 第 k 大的数即 nums 中索引为 k-1 的数，在舍弃区域比较值时都要 -1 处理
func findKth(nums1, nums2 []int, k int) int {
	n1, n2 := len(nums1), len(nums2)
	if n1 > n2 {
		n1, n2 = n2 , n1
		nums1, nums2 = nums2, nums1
	}
	if n1 == 0{
		return nums2[k-1]
	}
	if k == 1{
		return min(nums1[0], nums2[0])
	}
	k1 := min(k/2, n1)
	k2 := k - k1

	switch{
		case nums1[k1-1] < nums2[k2-1]:
			return findKth(nums1[k1:], nums2, k2)
		case nums1[k1-1] > nums2[k2-1]:
			return findKth(nums1, nums2[k2:], k1)
		default:
			return nums1[k1-1]
	}
}

func min(x, y int) int{
	if x > y{
		return y
	}
	return x
}
