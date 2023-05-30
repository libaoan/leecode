package main

// O(2n) 速度95% 内存5%
func moveZeroes(nums []int) {
	left, n := 0, len(nums)

	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[left] = nums[i]
			left++
		}

	}

	for left < n {
		nums[left] = 0
		left++
	}
}

// O(n) 速度25% 内存5%
func moveZeroes2(nums []int) {
	left, n := 0, len(nums)

	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}

	}

}
