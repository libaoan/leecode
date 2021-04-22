package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	//nums = []int{}
	//nums = []int{0}
	// res := threeSum(nums)
	sort.Ints(nums)
	fmt.Println("Before:", nums)
	length := len(nums)
	res := fourSum(nums, 0, length-1, 0)
	fmt.Println("After:", res)
}

// TODO 未调试
func fourSum(nums []int, left, right, target int) [][]int {

	var res [][]int
	length := len(nums)

	switch {
	case length < 4:
		return res
	case length == 4:
		{
			sum := nums[left] + nums[left+1] + nums[right-1] + nums[right]
			if sum == target {
				r := []int{nums[left], nums[left+1], nums[right-1], nums[right]}
				res = append(res, r)
				return res
			} else {
				return res
			}
		}
	case length > 4:
		{
			sum := nums[left] + nums[right]
			if sum >= target {
				right--
				return fourSum(nums[left:right+1], left, right, target)
			} else {

				left++
				right--
				twoTarget := target - (nums[left] + nums[right])
				resTwo2 := twoSum(nums[right:left], right, left, twoTarget)
				if resTwo2 != nil {
					ress := [][]int{}
					for _, x := range resTwo2 {
						resTwo := []int{nums[left], nums[right]}
						for _, n := range x {
							resTwo = append(resTwo, n)
						}
					}
					for _, r := range ress {
						res = append(res, r)
					}

				}
				right1 := right - 1
				res2 := fourSum(nums[left:right1+1], left, right1, target)
				if res2 != nil {
					for _, y := range res2 {
						res = append(res, y)
					}
				}

				left1 := left + 1
				res3 := fourSum(nums[left1:right+1], left1, right, target)
				if res3 != nil {
					for _, z := range res3 {
						res = append(res, z)
					}
				}
				return res

			}
		}
	}

	return res
}

func twoSum(nums []int, left, right, target int) [][]int {
	var res [][]int
	length := len(nums)

	switch {
	case length < 2:
		return res
	case length == 2:
		{
			sum := nums[left] + nums[right]
			if sum == target {
				r := []int{nums[left], nums[right]}
				res = append(res, r)
				return res
			} else {
				return res
			}
		}
	case length > 2:
		{
			sum := nums[left] + nums[right]
			if sum == target {
				r := []int{nums[left], nums[right]}
				res = append(res, r)
				left++
				right--
				res1 := twoSum(nums[left:right+1], left, right, target)
				if res1 != nil {
					for _, n := range res1 {
						res = append(res, n)
					}
				}
				return res
			}
			if sum > target {
				right--
				return twoSum(nums[left:right+1], left, right, target)
			} else {
				left++
				res2 := twoSum(nums[left:right+1], left, right, target)
				if res2 != nil {
					for _, x := range res2 {
						res = append(res, x)
					}
				}
				return res
			}
		}
	}

	return res
}
