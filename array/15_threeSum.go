package main

import (
	"fmt"
	"sort"
)

// todo 通过率 90% ， 超时
func threeSum(nums []int) [][]int {

	ans := make([][3]int, 0)
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				continue
			}
			sum := nums[i] + nums[j]
			for k, e := range nums {
				if e >= nums[j] && e == -sum && k != i && k != j {
					ans = append(ans, [3]int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	res := [][3]int{}
	for _, d := range ans {
		exist := false
		for _, c := range res {
			if d == c {
				exist = true
				break
			}
		}
		if !exist {
			res = append(res, d)
		}
	}
	ress := [][]int{}
	for i := 0; i < len(res); i++ {
		ress = append(ress, res[i][:])
	}
	return ress
}

func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(arr))
}
