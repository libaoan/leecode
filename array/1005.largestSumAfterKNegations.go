package main

import "sort"

func main() {

}

// todo： 待调试
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
			return sum(nums) - nums[0]
		}
	}

	if nums[cnt-1] == 0 {
		if cnt == n {
			return 0
		}
		return sum(nums[cnt:])
	}

	if cnt >= k {
		return sum(nums[k:])
	} else {
		if (k-cnt)%2 == 0 {
			return sum(nums[cnt:])
		} else if cnt == n {
			return 0
		} else {
			return sum(nums[cnt+1:])
		}
	}

}
