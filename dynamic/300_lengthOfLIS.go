package main

import (
	"fmt"
	"math"
	"time"
)

// todo 动态规划, 未实现
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	ans := make([][2]int, n)

	ans[0][0] = math.MaxInt
	ans[0][1] = nums[0]
	maxL := 1
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1]
		if nums[i] > ans[i-1][1] {
			ans[i][0] = ans[i-1][1]
			ans[i][1] = nums[i]
			maxL += 1
			continue
		}

		if nums[i] < ans[i-1][1] && nums[i] > ans[i-1][0] {
			ans[i][1] = nums[i]
		}
	}
	return maxL

}

// todo 单调栈， 失败，通过率50%
func lengthOfLIS2(nums []int) int {
	n, maxN := len(nums), 0

	for i := 0; i < n; i++ {
		var ans = []int{}
		for j := i; j < n; j++ {
			k := len(ans) - 1
			if k < 0 {
				ans = append(ans, nums[j])
				continue
			}

			if ans[k] < nums[j] {
				ans = append(ans, nums[j])
			} else {
				for k >= 0 && ans[k] >= nums[j] {
					ans = ans[:k]
					k--
				}
				ans = append(ans, nums[j])
			}
		}
		fmt.Println(ans)
		if maxN < len(ans) {
			maxN = len(ans)
		}

	}
	return maxN
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	nums = []int{0, 1, 0, 3, 2, 3}
	//nums = []int{3,5,6,2,5,4,19,5,6,7,12}
	start := time.Now().UnixNano()
	fmt.Println(lengthOfLIS2(nums))
	fmt.Println("运行时间", time.Now().UnixNano()-start)
}
