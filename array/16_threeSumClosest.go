package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 2, 1, -4}
	nums = []int{-86, 8, 0, -79, -23, 88, -26, -17, 41, 70, 1, 29, 3, -52, -30, -31, 14, 42, 4, -52, -77, -39, 91, -13, 98, 41, -16, -8, 91, 54, 50, 1, -22, 45, 30, 20, -36, -41, -48, -25, -3, 76, 69, -38, 78, -45, -75, -57, 55, -23, 65, -23, 69, 60, 85, -95, 84, -73, -21, 43, 35, 58, 66, 6, -49, 38, -5, 46, -24, 23, -23, 63, 55, -91, 80, -73, -53, 24, -77, -5, -93, 4, 2, -71, 23, 2, -44, -19, 33, 73, 9, -94, 41, 81, -14, 43, 27, 3, -65, -89, 0, -100, -89, -76, 95, 0, -60, 65, -77, 85, -29, 98, -8, 18, 41, -93, -64, 51, -77, -39, -75, -6, -65, 0, -15, 66, 80, 67, -96, 26, -57, -43, -76, 72, 20, -55, -83, 95, -65, 22, -47, -90, -83, 13, 6, 81, -23, -50, -17, 17, 100, 99, 66, -62, -21, 76, 37, 60, -88, 64, 78, -90, 67, -65, -62, 55, -87, -75, -79, 14, 87, -100, -87, 9, -98, 0, -83, 7, 29, 33, -87, -50, 6, -57, 28, 17, 26, 4}
	fmt.Println(threeSumClosest(nums, 136))
}

// todo 穷尽法， O(n3) 运行超时
func threeSumClosest(nums []int, target int) int {
	minus := []int{}
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if i != j && i != k && j != k {
					minus = append(minus, nums[i]+nums[j]+nums[k]-target)
				}
			}
		}
	}
	sort.Ints(minus)
	// fmt.Println(minus)
	i := 0
	for ; i < len(minus); i++ {
		if minus[i] >= 0 {
			break
		}
	}
	if i == len(minus) {
		return minus[i-1] + target
	}

	if i == 0 {
		return minus[0] + target
	}
	if 0-minus[i-1] < minus[i] {
		return minus[i-1] + target
	} else {
		return minus[i] + target
	}

}
