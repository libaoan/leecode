package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 5, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}

// 复杂度O(n3)， 速度 52%， 内存 70%
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
		p := nums[n-i]
		if !isMax(p, nums[n-i:]) {
			snums := make([]int, len(nums[n-i:]))
			copy(snums, nums[n-i:])
			sort.Ints(snums)
			for j := 0; j < len(snums); j++ {
				if snums[j] == p && snums[j+1] > p {
					k := j + 1
					nums = append(nums[:n-i], snums[k])
					nums = append(nums, snums[:k]...)
					if k < len(snums)-1 {
						nums = append(nums, snums[k+1:]...)
					}
					return
				}
			}
		}
	}
	sort.Ints(nums)

}
