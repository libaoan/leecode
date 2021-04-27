package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 0, -1, 0, -2, 2, -2}
	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	n := removeElement(nums, 2)
	fmt.Println("After:", n, nums)
}

// TODO: 未调试
func removeElement(nums []int, val int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	i := 0

	for {
		if (nums[i] - val) == 0 {

			if i != length-1 {
				for j, v := range nums[i:] {
					if v != val {
						nums[i] = v
						nums[j] = val
						break
					}
				}
				return i - 1
			}
			return i - 1
		} else {
			i++
		}
		if i == length {
			break
		}
	}
	return i
}
