package main

import "fmt"

func main() {
	nums := []int{3, 5, 2, 6}
	nums = []int{2, 4, 3, 3, 5, 4, 9, 6}
	// nums = []int{11,52,57,91,47,95,86,46,87,47,70,56,54,61,89,44,3,73,1,7,87,48,17,25,49,54,6,72,97,62,16,11,47,34,68,58,14,36,46,65,2,15}
	res := mostCompetitive3(nums, 4)
	for _, r := range res {
		fmt.Printf("%d ", r)
	}
}

// 基础算法，递归 时间复杂度 n!
func mostCompetitive(nums []int, k int) []int {

	if len(nums) < k || k <= 0 {
		return []int{}
	}

	if len(nums) == k {
		return nums
	}

	if k == 1 {
		min := nums[0]
		for _, num := range nums[1:] {
			switch {
			case num < min:
				return []int{num}
			case num > min:
				return []int{min}
			default:

			}
		}
		return []int{min}
	}

	bestInclude := []int{nums[0]}
	bestSub := mostCompetitive(nums[1:], k-1)
	bestInclude = append(bestInclude, bestSub...)

	bestExclude := mostCompetitive(nums[1:], k)

	for i := 0; i < len(bestExclude); i++ {
		switch {
		case bestInclude[i] > bestExclude[i]:
			return bestExclude
		case bestInclude[i] < bestExclude[i]:
			return bestInclude
		default:
		}
	}
	return bestInclude
}

// 基础算法优化，时间复杂度还是n!
func mostCompetitive2(nums []int, k int) []int {
	result := getAllSub(nums, k)
	if result == nil {
		return nil
	}
	minLi := result[0]
	for _, subLi := range result {
		for i := 0; i < len(minLi); i++ {
			if minLi[i] > subLi[i] {
				minLi = subLi
				break
			}
		}
	}
	return minLi

}

func getAllSub(nums []int, k int) [][]int {
	if len(nums) < k || k <= 0 {
		return nil
	}
	if len(nums) == k {
		return [][]int{nums}
	}
	if k == 1 {
		allSub := [][]int{}
		for _, num := range nums {
			allSub = append(allSub, []int{num})
		}
		return allSub
	}

	result := [][]int{}
	subLi := getAllSub(nums[1:], k-1)
	for _, li := range subLi {
		include := []int{nums[0]}
		include = append(include, li...)
		result = append(result, include)
	}
	excludeSub := getAllSub(nums[1:], k)
	if excludeSub != nil {
		result = append(result, excludeSub...)
	}
	return result
}

// 最佳答案，单调递增栈
func mostCompetitive3(nums []int, k int) []int {
	max_pop := len(nums) - k
	stack := []int{}

	for _, n := range nums {
		for len(stack) != 0 && stack[len(stack)-1] > n && max_pop > 0 {
			stack = stack[0 : len(stack)-1]
			max_pop--
		}
		stack = append(stack, n)
	}

	for max_pop > 0 {
		stack = stack[0 : len(stack)-1]
		max_pop--
	}
	return stack
}
