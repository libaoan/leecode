package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 3, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}

// 复杂度O(n3)， 通过率90%， 最后一步
func nextPermutation(nums []int) {

	// 1. 最后2位比较，如果次位小于最后一位，交换结果返回; 大于等于最后一位，检查倒数第3位
	// 2. 如果倒数第3位是后3位中最大的，检查倒数第4位，直到第1位
	// 3. 否则，如果倒数第3不是后3位中最大的，从小到大对后3位排序，其中倒数第3位选择次大的，返回

	n := len(nums)
	if n < 2 {
		return
	}

	isMax := func(x int, nums []int) bool {
		for _, a := range nums {
			if x < a {
				return false
			}
		}
		return true
	}

	if nums[n-2] < nums[n-1] {
		nums[n-2], nums[n-1] = nums[n-1], nums[n-2]
		return
	}

	for i := 3; i <= n; i++ {
		if !isMax(nums[n-i], nums[n-i:]) {
			p := nums[n-i]
			sort.Ints(nums[n-i:])
			for j, x := range nums[n-i:] {
				if x == p {
					nums[n-i], nums[n-i+j+1] = nums[n-i+j+1], nums[n-i]
					return
				}
			}
		}
	}
	sort.Ints(nums)

}
