package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	n := removeElement(nums, 2)
	fmt.Println(nums[:n])
}

// 双指针 O(n) 速度 100%, 内存 41%
func removeElement(nums []int, val int) int {

	i, j := 0, len(nums)-1

	for i <= j {
		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else {
			i++
		}
	}
	return i
}
