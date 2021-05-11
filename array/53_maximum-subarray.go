package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//nums = []int{-10000}
	res := maxSubArray(nums)
	res = maxSubArray2(nums)

	fmt.Printf("the res:", res)
}

// 暴力法
func maxSubArray(nums []int) int {
	length := len(nums)
	max := nums[0]
	for i := 0; i < length; i++ {
		for j := 1; j <= length-i; j++ {
			sum := sumSlice(nums[i : i+j])
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

func sumSlice(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

// 贪心算法
func maxSubArray2(nums []int) int {
	length := len(nums)
	sum := nums[0]
	max := nums[0]
	for i := 1; i < length; i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum = sum + nums[i]
		}

		if max < sum {
			max = sum
		}
	}
	return max
}

// 动态规划算法
func maxSubArray3(nums []int) int {
	max := nums[0]
	length := len(nums)
	for i := 1; i < length; i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}

		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// TODO: 分治算法
func maxSubArray4(nums []int) int {
	return 0
}
