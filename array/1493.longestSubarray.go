package main

func main() {

}

// 参考滑动窗口 + 窗口大小二分法， 通过率 90%, todo：运行超时
func longestSubarray(nums []int) int {
	start, end := 0, len(nums)
	ans := -1
	for start <= end {
		// 划窗大小
		win := start + (end-start)/2
		if Judge(nums, win) {
			ans, start = win, win+1
		} else {
			end = end - 1
		}
	}

	if ans > 0 {
		return ans - 1
	} else {
		return -1
	}
}

func Judge(nums []int, win int) bool {
	numCnts := [2]int{}
	for i := 0; i < win; i++ {
		numCnts[nums[i]]++
	}
	if numCnts[0] <= 1 {
		return true
	}
	for i := win; i < len(nums); i++ {
		numCnts[nums[i-win]]--
		numCnts[nums[i]]++
		if numCnts[0] <= 1 {
			return true
		}
	}
	return false
}
