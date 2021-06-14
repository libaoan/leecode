package main

import "fmt"

func main() {

	nums := []int{5, 1, 6}
	//nums = []int{1,3}
	nums = []int{3, 4, 5, 6, 7, 8}
	fmt.Println(subsetXORSum(nums))
}

func subsetXORSum(nums []int) int {
	ln := len(nums)
	if ln == 0 {
		return 0
	}

	if ln == 1 {
		return nums[0]
	}

	if ln == 2 {
		return nums[0] ^ nums[1] + nums[0] + nums[1]
	}

	res := nums[0]
	for _, x := range nums[1:] {
		res ^= x
	}

	x := nums[0]
	newNums := nums[1:]
	sum := x + subsetXORSum(newNums)

	for i, _ := range nums {
		subNums := []int{}
		if i == 0 {
			subNums = nums[1:]
		} else if i == ln-1 {
			subNums = nums[:ln-1]
		} else {
			subNums = append(nums[:i], nums[i+1:]...)
		}
		t := x
		for _, y := range subNums {
			t ^= y
		}
		sum += t
	}

	return res + sum
}
