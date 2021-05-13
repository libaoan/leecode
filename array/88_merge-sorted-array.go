package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	nums1 = []int{0, 0, 0, 0}
	merge4(nums1, 0, nums2, 3)
	fmt.Println(nums1)
}

// 双指针
func merge2(nums1 []int, m int, nums2 []int, n int) {

	newNums := make([]int, m+n)
	i, j, k := 0, 0, 0
	for {
		if i >= m || j < n && nums1[i] >= nums2[j] {
			newNums[k] = nums2[j]
			j++
			k++
		} else {
			newNums[k] = nums1[i]
			i++
			k++
		}

		if k == m+n {
			break
		}
	}
	for i, v := range newNums {
		nums1[i] = v
	}
}

// 双指针法优化
func merge21(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	copy(nums1, sorted)
}

// 采用内建函数，把nums2加入到num1尾部，然后进行排序
func merge3(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

// 逆指针算法， 节省空间
func merge4(nums1 []int, m int, nums2 []int, n int) {

	p1, p2 := 0, 0
	for {

		if p1 >= m {
			copy(nums1[0:n-p2], nums2[0:n-p2])
			break
		}

		if p2 >= n {
			break
		}

		if nums1[m-p1-1] < nums2[n-p2-1] {
			nums1[m+n-p1-p2-1] = nums2[n-p2-1]
			p2++
		} else {
			nums1[m+n-p1-p2-1] = nums1[m-p1-1]
			p1++
		}

	}
}
