package main

// T(O) = 83% S(O) = 86%
func sortedSquares(nums []int) []int {
	le, ri, ln := -1, -1, len(nums)
	for i := 0; i < ln; i++ {
		if nums[i] >= 0 {
			ri = i
			break
		}
	}
	if ri < 0 {
		le = ln - 1
	}
	if ri >= 0 {
		le = ri - 1
	}

	res := []int{}
	for le >= 0 && ri >= 0 && ri < ln {
		if -nums[le] < nums[ri] {
			res = append(res, nums[le]*nums[le])
			le--
		} else {
			res = append(res, nums[ri]*nums[ri])
			ri++
		}
	}
	for le >= 0 {
		res = append(res, nums[le]*nums[le])
		le--
	}
	for ri >= 0 && ri < ln {
		res = append(res, nums[ri]*nums[ri])
		ri++
	}
	return res

}
