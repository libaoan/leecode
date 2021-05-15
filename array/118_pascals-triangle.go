package main

import "fmt"

func main() {
	nums := generate(8)
	for _, num := range nums {
		for _, n := range num {
			fmt.Print(n, ",")
		}
		fmt.Println()
	}
}

// 基础算法
func generate(numRows int) [][]int {

	var nums [][]int
	for i := 0; i < numRows; i++ {
		if i == 0 {
			nums = append(nums, []int{1})
		} else {
			newNums := make([]int, i/2+1)
			newNums[0] = 1
			if (i+1)%2 == 0 {
				mid := i / 2
				for j := 1; j <= mid; j++ {
					newNums[j] = nums[i-1][j-1] + nums[i-1][j]
				}
				for j := mid; j >= 0; j-- {
					newNums = append(newNums, newNums[j])
				}
			} else {
				mid := i / 2
				for j := 1; j <= mid; j++ {
					newNums[j] = nums[i-1][j-1] + nums[i-1][j]
				}
				for j := mid - 1; j >= 0; j-- {
					newNums = append(newNums, newNums[j])
				}
			}
			nums = append(nums, newNums)
		}
	}
	return nums
}
