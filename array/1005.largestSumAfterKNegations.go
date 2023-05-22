package main

import (
	"fmt"
	"sort"
)

func main() {
	nums, k := []int{4, 3, 2}, 1
	nums, k = []int{3, -1, 0, 2}, 3
	nums, k = []int{2, -3, -1, 5, -4}, 2
	nums, k = []int{5, 6, 9, -3, 3}, 2
	nums, k = []int{-2, 9, 9, 8, 4}, 5
	nums, k = []int{-8, 3, -5, -3, -5, -2}, 6
	fmt.Println(largestSumAfterKNegations(nums, k))
}

// todo：通过率99%
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)

	sum := func(nums []int) int {
		res := 0
		for _, c := range nums {
			res += c
		}
		return res
	}

	cnt, n := 0, len(nums)

	for _, c := range nums {
		if c <= 0 {
			cnt++
		} else {
			break
		}
	}

	if cnt == 0 {
		if k%2 == 0 {
			return sum(nums)
		} else {
			return sum(nums[1:]) - nums[0]
		}
	}

	if nums[cnt-1] == 0 {
		if cnt >= k {
			return sum(nums[cnt:]) - sum(nums[:k]) + sum(nums[k:cnt])
		} else {
			return sum(nums[cnt:]) - sum(nums[:cnt])
		}
	}

	if cnt >= k {
		return sum(nums[cnt:]) - sum(nums[:k]) + sum(nums[k:cnt])
	} else {
		if (k-cnt)%2 == 0 {
			return sum(nums[cnt:]) - sum(nums[:cnt])
		} else if cnt == n {
			return -sum(nums)
		} else {
			return sum(nums[cnt+1:]) - sum(nums[:cnt+1])
		}
	}

}
