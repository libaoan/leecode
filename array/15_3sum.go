package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	//nums = []int{}
	//nums = []int{0}
	// res := threeSum(nums)
	res := threeSum2(nums)
	fmt.Println(res)
}

//暴力法
func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
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

// TODO 优化， 有序序列两端向中间遍历
func threeSum2(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	fRange(0, len(nums)-1, nums, res)
	return res
}

func fRange(left, right int, nums []int, res [][]int) {
	if left >= right {
		return
	}
	for _, x := range nums[left+1 : right] {
		sum := nums[left] + nums[right] + x
		switch {

		case sum == 0:
			{
				r := []int{nums[left], x, nums[right]}
				res = append(res, r)
			}
		case sum < 0:
			{
				left++
				fRange(left, right, nums, res)
			}
		case sum > 0:
			{
				right--
				fRange(left, right, nums, res)
			}

		}

	}
}
