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
	nums, k = []int{8, -7, -3, -9, 1, 9, -6, -9, 3}, 8
	nums, k = []int{8, -7, -1, -9, 2, 9, -6, -9, 3}, 8
	nums, k = []int{2, -3, -1, 5, -4}, 2
	fmt.Println(largestSumAfterKNegations3(nums, k))
}

// O(n2) 速度 80%, 内存80%
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
			return -sum(nums[:cnt]) + 2*nums[cnt-1]
		} else {
			//fmt.Println(nums)
			if -nums[cnt-1] < nums[cnt] {
				return sum(nums[cnt:]) - sum(nums[:cnt]) + 2*nums[cnt-1]
			} else {
				return sum(nums[cnt:]) - sum(nums[:cnt]) - 2*nums[cnt]
			}

		}
	}

}

// todo： 参考别人的答案， 贪心算法 O(n) 速度 80%, 内存3%
func largestSumAfterKNegations2(nums []int, k int) (ans int) {

	maps := make(map[int]int, 0)

	for _, c := range nums {
		maps[c]++
		ans += c
	}

	for i := -100; i < 0 && k > 0; i++ {
		if maps[i] > 0 {
			ops := min(k, maps[i])
			ans -= i * ops * 2
			maps[-i] += ops
			k -= ops
		}
	}

	if maps[0] == 0 && k > 0 && k%2 != 0 {
		for i := 1; i <= 100; i++ {
			if maps[i] > 0 {
				ans -= i * 2
				break
			}
		}
	}

	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// todo： 参考别人的答案， 贪心算法 O(n) 速度 86%, 内存80%
func largestSumAfterKNegations3(nums []int, k int) (ans int) {

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		// 贪心：如果是负数，而k还有盈余，就把负数反过来
		if nums[i] < 0 && k > 0 {
			nums[i] = -1 * nums[i]
			k--
		}
		ans += nums[i]
	}

	sort.Ints(nums)

	if k%2 != 0 {
		ans -= 2 * nums[0]
	}

	return
}
