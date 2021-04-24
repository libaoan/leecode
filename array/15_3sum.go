package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	//nums = []int{}
	//nums = []int{0}
	nums = []int{-9, -8, -2, -1, 0, 1, 3, 5, 9}
	nums = []int{-9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	sort.Ints(nums)
	fmt.Println("Before:", nums)
	res := threeSum(nums)
	fmt.Println("After:", res)
	res2 := threeSum2(nums)
	fmt.Println("After:", res2)
}

//暴力法
func threeSum(nums []int) [][]int {
	res := [][]int{}
	length := len(nums)
	for i := 0; i < length; i++ {
		minus := 0 - nums[i]
		for j := i + 1; j < length; j++ {
			n := minus - nums[j]
			if Contains(nums[j+1:], n) {
				item := []int{nums[i], nums[j], n}
				if !SliceContains(res, item) {
					res = append(res, item)
				}
			}
		}
	}
	return res
}

func Contains(li []int, n int) bool {
	for _, x := range li {
		if x == n {
			return true
		}
	}
	return false
}

func SliceContains(li [][]int, item []int) bool {
	for _, x := range li {
		if len(x) != len(item) {
			return false
		}
		flag := false
		for i, y := range x {
			if y != item[i] {
				flag = true
				break
			}
		}

		if !flag {
			return true
		}
	}

	if len(item) == 0 && len(li) == 0 {
		return true
	} else {
		return false
	}
}

// 优化1， 有序序列两端向中间遍历
func threeSum2(nums []int) [][3]int {
	var res [][3]int
	length := len(nums)
	max := length - 1
	switch {
	case length < 3:
		{
			return res
		}
	case length == 3:
		{
			sum := nums[0] + nums[1] + nums[2]
			if sum == 0 {
				r := [3]int{nums[0], nums[1], nums[2]}
				res = append(res, r)
			}
		}
	case length > 3:
		{

			part := 0 - (nums[0] + nums[max])
			if Contains(nums[1:max], part) {
				r := [3]int{nums[0], part, nums[max]}
				res = append(res, r)
			}
			// 2、 加入最小值，排除最大值
			sNums := nums[1:max]
			ress1 := twoSums(sNums, 0-nums[0])
			if ress1 != nil {
				for _, item := range ress1 {
					r := [3]int{nums[0], item[0], item[1]}
					res = append(res, r)
				}
			}
			// 3、 加入最大值，排除最小值
			sNums = nums[1:max]
			ress2 := twoSums(sNums, 0-nums[max])
			if ress2 != nil {
				for _, item := range ress2 {
					r := [3]int{item[0], item[1], nums[max]}
					res = append(res, r)
				}
			}
			// 4、最后在递归子序列
			sNums = nums[1:max]
			ress := threeSum2(sNums)
			if ress != nil {
				for _, item := range ress {
					res = append(res, item)
				}
			}

		}
	}
	return res
}

func twoSums(nums []int, target int) [][2]int {
	var res [][2]int
	length := len(nums)
	max := length - 1
	switch {
	case length < 2:
		{
			return res
		}
	case length == 2:
		{
			if (nums[0] + nums[1]) == target {
				r := [2]int{nums[0], nums[1]}
				res = append(res, r)
			}
		}
	case length > 2:
		{
			// 第一种、两端都能加进去
			part := target - (nums[0] + nums[max])
			if part == 0 {
				//1、两端加进去
				r := [2]int{nums[0], nums[max]}
				res = append(res, r)

				//2、 加入最小值，排除最大值，寻找另一个值
				sNums := nums[1:max]
				item := target - nums[0]
				if Contains(sNums, item) {
					r := [2]int{nums[0], item}
					res = append(res, r)
				}

				//3、 加入最大值，排除最小值，寻找另一个值
				sNums = nums[1:max]
				item = target - nums[max]
				if Contains(sNums, item) {
					r := [2]int{item, nums[max]}
					res = append(res, r)
				}

				//4、递归子序列
				sNums = nums[1:max]
				ress := twoSums(sNums, target)
				if ress != nil {
					for _, item := range ress {
						res = append(res, item)
					}
				}

			}
			// 第二种:只加最左端（最小），排除最右端
			if part < 0 {
				sNums := nums[1:max]
				item := target - nums[0]
				if Contains(sNums, item) {
					r := [2]int{nums[0], item}
					res = append(res, r)
				}
				sNums = nums[1:max]
				ress := twoSums(sNums, target)
				if ress != nil {
					for _, item := range ress {
						res = append(res, item)
					}
				}
			}
			//  第三种:只加最右端（最大），排除最左端
			if part > 0 {
				sNums := nums[1:max]
				item := target - nums[max]
				if Contains(sNums, item) {
					r := [2]int{item, nums[max]}
					res = append(res, r)
				}
				sNums = nums[1:max]
				ress := twoSums(sNums, target)
				if ress != nil {
					for _, item := range ress {
						res = append(res, item)
					}
				}
			}

		}
	}
	return res
}
