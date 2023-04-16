package main

import (
	"fmt"
)

// 动态规划算法
// fn = 1 n=1,2,4,8 ...
// fn = f(n-1) + 1 n%2 != 0
// fn = f(n-1)     n%2 == 0
// [todo] f(12) + 不符合
func countBits(n int) []int {

	ans := make([]int, n+1)
	ans[0] = 0
	for i := 1; i <= n; i++ {
		b := 1
		for ; b < i; b = b << 1 {
			print(b, " ")
		}
		println()
		switch {
		case b == i:
			ans[i] = 1
		case i%2 == 0:
			ans[i] = ans[i-1]
		default:
			ans[i] = ans[i-1] + 1
		}
	}
	return ans
}

func main() {
	fmt.Println(countBits(12))
}
