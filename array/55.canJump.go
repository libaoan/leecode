package main

import "fmt"

func main() {

	nums := []int{2, 3, 1, 1, 4}
	nums = []int{3, 2, 1, 0, 4}
	nums = []int{1, 1, 2, 2, 0, 1, 1}
	nums = []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}

	fmt.Println(canJump2(nums))
}

// O(n3) todo 通过率 90%
func canJump(nums []int) bool {

	n := len(nums) - 1
	if n < 1 {
		return true
	}

	for start := nums[0]; start >= 1; start-- {
		curStep := start
		if curStep >= n {
			return true
		}
		nextStep := curStep + nums[curStep]
		for nextStep > curStep {
			for nextStep < n && nums[nextStep] != 0 {
				curStep = nextStep
				nextStep = curStep + nums[curStep]
				//nextStep = nextStep + nums[nextStep]
			}

			if nextStep >= n {
				return true
			} else {
				nextStep--
			}

			fmt.Println(curStep, nums[curStep], nextStep, nums[nextStep])
		}

	}
	return false
}

// 参考答案: todo 贪心算法
func canJump2(nums []int) bool {
	cover := 0
	n := len(nums) - 1
	for i := 0; i <= cover; i++ { // 每次与覆盖值比较
		cover = max(i+nums[i], cover) //每走一步都将 cover 更新为最大值
		if cover >= n {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
