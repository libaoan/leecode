package main

import "fmt"

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))

	fmt.Println()

	fmt.Println(containsNearbyDuplicate2([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicate2([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicate2([]int{1, 2, 3, 1, 2, 3}, 2))
}

// 哈希查找法， 空间复杂度n， 时间复杂度n
func containsNearbyDuplicate(nums []int, k int) bool {
	maps := make(map[int]int)

	for i := range nums {
		if v, ok := maps[nums[i]]; !ok {
			maps[nums[i]] = i
		} else {
			if v+k >= i {
				return true
			} else {
				maps[nums[i]] = i
			}
		}
	}

	return false
}

// 线性查找法 n*n
func containsNearbyDuplicate2(nums []int, k int) bool {

	for i := range nums {

		for j := i + 1; j < len(nums) && j <= i+k; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

// todo: 以上还不是最优解，待优化
func containsNearbyDuplicate3(nums []int, k int) bool {
	return false
}
