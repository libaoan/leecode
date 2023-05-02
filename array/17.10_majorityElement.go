package main

import "fmt"

func main() {
	nums := []int{1, 2, 5, 9, 5, 9, 5, 5, 5}
	nums = []int{2, 2, 1, 1, 1, 2, 2}
	nums = []int{6, 5, 5}
	fmt.Println(majorityElement(nums))
}

// todo 待调试
func majorityElement(nums []int) int {
	n := len(nums)
	i, j := 1, n-1
	cnt := 1
	for i <= j {
		if nums[i] < nums[0] {
			i++
		}
		if nums[i] > nums[0] {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
		if nums[i] == nums[0] {
			nums[i] = -1
			i++
			cnt++
		}
	}

	// fmt.Println(nums, cnt, i, j)

	if cnt > n/2 {
		return nums[0]
	}

	if n-j > n/2 {
		return majorityElement(nums[j+1:])
	}
	k := 0
	for _, x := range nums[1:i] {
		if x != -1 {
			nums[k] = x
			k++
		}
	}
	if len(nums[:k]) > n/2 {
		return majorityElement(nums[:k])
	}

	return -1
}
