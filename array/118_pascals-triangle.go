package main

import "fmt"

func main() {
	nums := generate(8)
	nums = generate2(8)
	for _, num := range nums {
		for _, n := range num {
			fmt.Print(n, ",")
		}
		fmt.Println()
	}
}

//  基础算法, 代码优化
func generate(numRows int) [][]int {

	var nums [][]int
	for i := 0; i < numRows; i++ {
		if i == 0 {
			nums = append(nums, []int{1})
		} else {
			newNums := make([]int, i/2+1)
			newNums[0] = 1
			mid := i / 2
			for j := 1; j <= mid; j++ {
				newNums[j] = nums[i-1][j-1] + nums[i-1][j]
			}
			if i%2 == 0 {
				mid -= 1
			}
			for j := mid; j >= 0; j-- {
				newNums = append(newNums, newNums[j])
			}
			nums = append(nums, newNums)
		}
	}
	return nums
}

// 别人家的代码， 算法复杂度相同，但是更代码优雅
func generate2(numRows int) [][]int {
	nums := make([][]int, numRows)
	for i := range nums {
		nums[i] = make([]int, i+1)
		nums[i][0] = 1
		nums[i][i] = 1
		for j := 1; j < i; j++ {
			nums[i][j] = nums[i-1][j-1] + nums[i-1][j]
		}
	}
	return nums
}
