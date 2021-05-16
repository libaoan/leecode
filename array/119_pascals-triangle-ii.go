package main

import "fmt"

func main() {
	nums := getRow(8)
	nums = getRow2(8)
	for _, n := range nums {
		fmt.Print(n, ",")
	}

}

// 基础算法
func getRow(rowIndex int) []int {
	if rowIndex < 0 {
		return nil
	}
	nums := make([][]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		nums[i] = make([]int, i+1)
		nums[i][0] = 1
		nums[i][i] = 1
		for j := 1; j < i; j++ {
			nums[i][j] = nums[i-1][j-1] + nums[i-1][j]
		}
	}
	return nums[rowIndex]
}

// 优化空间复杂度
func getRow2(rowIndex int) []int {
	if rowIndex < 0 {
		return nil
	}
	var pre, cur []int
	for i := 0; i <= rowIndex; i++ {
		cur = make([]int, i+1)
		cur[0] = 1
		cur[i] = 1
		for j := 1; j < i; j++ {
			cur[j] = pre[j-1] + pre[j]
		}
		pre = cur
	}
	return cur
}
