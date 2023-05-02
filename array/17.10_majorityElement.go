package main

import "fmt"

func main() {
	nums := []int{1, 2, 5, 9, 5, 9, 5, 5, 5}
	nums = []int{2, 2, 1, 1, 1, 2, 2}
	nums = []int{6, 5, 5}
	nums = []int{2, 2, 2, 3, 3, 4, 4}
	fmt.Println(majorityElement(nums))
}

// 速度 100%， 内存63%
func majorityElement(nums []int) int {

	return reg(nums, len(nums))

}

func reg(nums []int, nl int) int {
	n := len(nums)
	i, j := 1, n-1
	cnt := 1
	for i <= j {
		if nums[i] < nums[0] {
			i++
		} else if nums[i] > nums[0] {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else {
			nums[i] = -1
			i++
			cnt++
		}
	}

	// fmt.Println(nums, cnt, i, j)

	if cnt > nl/2 {
		return nums[0]
	}

	if n-j > nl/2 {
		return reg(nums[j+1:], nl)
	}
	k := 0
	for _, x := range nums[1:i] {
		if x != -1 {
			nums[k] = x
			k++
		}
	}
	if len(nums[:k]) > n/2 {
		return reg(nums[:k], nl)
	}

	return -1
}
