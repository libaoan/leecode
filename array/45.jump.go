package main

import "fmt"

func main() {

	nums := []int{2, 3, 1, 1, 4}
	nums = []int{1, 1, 1, 1}
	nums = []int{1, 2, 1, 1, 1}
	nums = []int{4, 1, 1, 3, 1, 1, 1}
	fmt.Println(jump(nums))

}

// O(n) 速度80% 内存50%
func jump(nums []int) int {

	max := func(i, j int, nums []int) int {
		for k := i + 1; k <= j; k++ {
			if k+nums[k] > i+nums[i] {
				i = k
			}
		}
		return i
	}

	i, jmp, n := 0, 1, len(nums)-1

	if n < 1 {
		return 0
	}

	for i+nums[i] < n {
		jmp++
		step := max(i, i+nums[i], nums)
		if step == i {
			i += nums[i]
		} else {
			i = step
		}
		println("-", i)
	}
	return jmp
}
