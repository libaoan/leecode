package main

// O(N) 速度96%， 内存62%
func averageValue(nums []int) int {
	sum, cnt, n := 0, 0, len(nums)
	for i := 0; i < n; i++ {
		if nums[i]%6 == 0 {
			sum += nums[i]
			cnt++
		}
	}
	if cnt == 0 {
		return 0
	} else {
		return sum / cnt
	}

}
