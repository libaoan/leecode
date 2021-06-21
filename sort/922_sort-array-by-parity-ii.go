package main

import "fmt"

func main() {
	nums := []int{4, 2, 5, 7}
	nums1 := sortArrayByParityII3(nums)
	fmt.Println(nums1)
}

func sortArrayByParityII(nums []int) []int {
	qi := []int{}
	ou := []int{}

	for _, x := range nums {
		if x%2 == 0 {
			ou = append(ou, x)
		} else {
			qi = append(qi, x)
		}
	}

	q, o := 0, 0
	for i := range nums {
		if i%2 == 0 {
			nums[i] = ou[o]
			o++
		} else {
			nums[i] = qi[q]
			q++
		}
	}
	return nums
}

// 和第一个算法一样
func sortArrayByParityII2(a []int) []int {
	ans := make([]int, len(a))
	i := 0
	for _, v := range a {
		if v%2 == 0 {
			ans[i] = v
			i += 2
		}
	}
	i = 1
	for _, v := range a {
		if v%2 == 1 {
			ans[i] = v
			i += 2
		}
	}
	return ans
}

// 双指针算法， 时间复杂度n， 空间复杂度1
func sortArrayByParityII3(nums []int) []int {
	o, q := 0, 1
	for q < len(nums) && o < len(nums) {
		if nums[o]%2 == 0 {
			o += 2
		} else {
			nums[o], nums[q] = nums[q], nums[o]
			q += 2
		}

	}

	return nums
}
