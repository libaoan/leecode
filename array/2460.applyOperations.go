package main

// O(n)， 速度 60%， 内存100%
func applyOperations(nums []int) []int {
	cnt0, n := 0, len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] == 0 {
			cnt0++
		} else if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}

		if i == n-2 && nums[n-1] == 0 {
			cnt0++
		}
	}
	j := 0
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for j < n {
		nums[j] = 0
		j++
	}
	return nums
}
