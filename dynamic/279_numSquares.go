package main

import (
	"fmt"
	"time"
)

// 递归算法， 结果正确， 运算超时，强制要求用动态规划？
func numSquares2(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}

	var cnts []int
	for i := n; i > 0; i-- {
		minus := n - i*i
		if minus == 0 {
			return 1
		} else if minus > 0 {
			cnts = append(cnts, numSquares(minus)+1)
		}

	}

	ans := cnts[0]
	for _, cnt := range cnts {
		if ans > cnt {
			ans = cnt
		}
	}

	return ans
}

func main() {
	start := time.Now().UnixNano()
	fmt.Println(numSquares2(58))
	fmt.Println("运行时间", time.Now().UnixNano()-start)
}
