package main

import "fmt"

// 算法负责度从n3 下降到n2， 利用分治算法实现
func countTriplets(nums []int) (ans int) {

	cnt := [1 << 16]int{}

	// 计算2个数按位与的结果统计
	for _, i := range nums {
		for _, j := range nums {
			cnt[i&j]++
		}
	}

	// 遍历2位与结果统计数组，与第3位与运算的结果是0的计数
	for x, c := range cnt {
		for _, y := range nums {
			if x&y == 0 {
				ans += c
			}
		}
	}

	return
}

func main() {

	fmt.Println(countTriplets([]int{2, 1, 3}))
}
