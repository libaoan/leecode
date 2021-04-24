package main

import "fmt"

func main() {
	var nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	n := removeDuplicates(nums)
	fmt.Println(n, nums[0:n])
}

func removeDuplicates(nums []int) int {

	length := len(nums)
	if length < 2 {
		return length
	}
	j := 0
	for i := 0; i < length; i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1
}
