package main

import (
	"fmt"
	"time"
)

// 动态规划 速度 12%， 空间 30%， todo 待优化
func massage(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n < 2 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]

	max := func(i, j int) int {
		if i < j {
			return j
		} else {
			return i
		}
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if j+1 != i {
				dp[i] = max(dp[i], dp[j]+nums[i])
			} else if j != 0 {
				dp[i] = max(dp[i], dp[j-1]+nums[i])
			} else {
				dp[i] = nums[i]
			}
		}
		fmt.Println(dp)
	}

	res := 0
	for _, r := range dp {
		if res < r {
			res = r
		}
	}
	return res
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	nums = []int{0, 1, 0, 3, 2, 3}
	//nums = []int{3,5,6,2,5,4,19,5,6,7,12}
	start := time.Now().UnixNano()
	fmt.Println(massage(nums))
	fmt.Println("运行时间", time.Now().UnixNano()-start)
}
