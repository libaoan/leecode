package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(numbers []int, target int) []int {

	biSearch := func(orderedList []int, target, index int) int {

		target = target - orderedList[index]
		low, hi := 0, len(orderedList)-1

		for hi >= low {
			mid := low + (hi-low)/2
			switch {
			case orderedList[mid] < target:
				low = mid + 1
			case orderedList[mid] > target:
				hi = mid - 1
			case orderedList[mid] == target && mid != index:
				return mid
			case orderedList[mid] == target && mid == index:
				if mid < len(orderedList)-1 && orderedList[mid+1] == target {
					return mid + 1
				}
				if mid > 0 && orderedList[mid-1] == target {
					return mid - 1
				} else {
					return -1
				}
			}
		}
		return -1
	}

	for index := range numbers {
		oIndex := biSearch(numbers, target, index)
		if oIndex != -1 {
			return []int{index, oIndex}
		}
	}

	return nil
}
