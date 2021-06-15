package main

import "fmt"

func main() {

	nums := []int{5, 1, 6}
	//nums = []int{1,3}
	nums = []int{3, 4, 5, 6, 7, 8}
	fmt.Println(subsetXORSum(nums))
}

// 基础迭代算法
func subsetXORSum(nums []int) int {
	ln := len(nums)
	if ln == 0 {
		return 0
	}

	if ln == 1 {
		return nums[0]
	}

	if ln == 2 {
		return nums[0] ^ nums[1] + nums[0] + nums[1]
	}

	sum := 0
	// 计算包含X元素的集合之异或之和
	for i := 0; i < ln; i++ {

		subs := getSubset(nums[1:], i)
		if subs == nil {
			sum = nums[0]
		}

		for _, sub := range subs {

			x := nums[0]
			for _, s := range sub {
				x ^= s
			}
			sum += x

		}
	}

	// 计算不包含X的真子集异或之和
	sum += subsetXORSum(nums[1:])

	return sum
}

func getSubset(nums []int, n int) [][]int {
	if n == 0 || n > len(nums) {
		return nil
	}
	if n == 1 {
		subset := [][]int{}
		for i := range nums {
			subset = append(subset, []int{nums[i]})
		}
		return subset
	}

	if n == len(nums) {
		return [][]int{nums}
	}

	subset := [][]int{}
	// Include nums[0] and chose n-1 in nums[1:]
	x := nums[0]
	subs := getSubset(nums[1:], n-1)
	if subs == nil {
		subset = append(subset, []int{x})
	}
	for _, sub := range subs {
		subWithX := []int{x}
		subWithX = append(subWithX, sub...)
		subset = append(subset, subWithX)
	}
	// Exclude nums[0] and chose n in nums[1:]
	subs = getSubset(nums[1:], n)
	if subs != nil {
		subset = append(subset, subs...)
	}

	return subset
}

// TODO: 优化算法待实现
func subsetXORSum2(nums []int) int {
	return 0
}
