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
				newNums := nums[left : right+1]
				return fourSum(newNums, 0, len(newNums)-1, target)
			} else {

				twoLeft := left + 1
				twoRight := right - 1
				twoTarget := target - (nums[left] + nums[right])
				twoNums := nums[twoLeft : twoRight+1]
				resTwo2 := twoSum(twoNums, 0, len(twoNums)-1, twoTarget)
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
				newNums := nums[left:right]
				res2 := fourSum(newNums, 0, len(newNums)-1, target)
				if res2 != nil {
					for _, y := range res2 {
						res = append(res, y)
					}
				}

				newNums = nums[left+1 : right+1]
				res3 := fourSum(newNums, 0, len(newNums)-1, target)
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
				newNums := nums[left+1 : right]
				res1 := twoSum(newNums, 0, len(newNums)-1, target)
				if res1 != nil {
					for _, n := range res1 {
						res = append(res, n)
					}
				}
				return res
			}
			if sum > target {
				newNums := nums[left:right]
				return twoSum(newNums, 0, len(newNums)-1, target)
			} else {

				newNums := nums[left+1 : right+1]
				res2 := twoSum(newNums, 0, len(newNums)-1, target)
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
