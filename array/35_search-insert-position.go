package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	//index := searchInsert(nums, 5)
	index := searchInsert2(nums, 7)
	fmt.Printf("index of the target %v", index)
}

// 遍历查找
func searchInsert(nums []int, target int) int {
	length := len(nums)
	index := 0
	for i := 0; i < length; i++ {
		if nums[i] == target {
			return i
		}
		if nums[i] > target {
			index = i
			break
		} else {
			index++
		}
	}
	return index
}

// 二分查找
func searchInsert2(nums []int, target int) int {
	length := len(nums)
	index := length
	left, right := 0, length-1

	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			index = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return index
}
