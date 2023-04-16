package main

import (
	"fmt"
)

// 动态规划算法， 时间 85%, 空间 99%
// fn = 1 n=1,2,4,8 ...
// fn = f(n) + f(n-x)
func countBits(n int) []int {

	ans := make([]int, n+1)
	ans[0] = 0
	for i := 1; i <= n; i++ {
		b := 1
		for ; b < i; b = b << 1 {
			//print(b, " ")
		}
		// println()
		if b == i {
			ans[i] = 1
		} else {
			b = b >> 1
			ans[i] = ans[b] + ans[i-b]
		}
	}
	return ans
}

func main() {
	fmt.Println(countBits(12))
}
