package main

import "fmt"

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}

	l1 := nextGreaterElement2(nums1, nums2)
	for _, x := range l1 {
		fmt.Print(x, " ")
	}
}

// 最基础做法，时间复杂度O(N*N)不符合要求
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	newNums := []int{}

	for _, x := range nums1 {
		flag := false
		for _, y := range nums2 {
			switch {
			case x == y:
				flag = true
			case flag == true && y > x:
				flag = false
				newNums = append(newNums, y)
			default:
			}
		}
		if flag {
			newNums = append(newNums, -1)
		}
	}
	return newNums
}

//采用单调栈解决
func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	stack := []int{}
	maps := make(map[int]int)
	for _, x := range nums2 {
		if len(stack) == 0 || stack[len(stack)-1] >= x {
			stack = append(stack, x)
		} else {
			for len(stack) != 0 && stack[len(stack)-1] < x {
				maps[stack[len(stack)-1]] = x
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, x)
		}
	}
	for len(stack) != 0 {
		maps[stack[len(stack)-1]] = -1
		stack = stack[:len(stack)-1]
	}

	for i, x := range nums1 {
		nums1[i] = maps[x]
	}
	return nums1
}
