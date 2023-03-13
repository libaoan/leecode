package main

import (
	"fmt"
	"math"
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
			cnts = append(cnts, numSquares2(minus)+1)
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

// 动态规划实现 速度43%， 内存92%  有优化空间， 复杂度n2
func numSquares(n int) int {

	min := func(i, j int) int {
		if i <= j {
			return i
		}
		return j
	}

	cnts := make([]int, n+1)
	for i := 0; i <= n; i++ {
		cnts[i] = math.MaxInt
	}
	cnts[0] = 0
	cnts[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; i >= j*j; j++ {
			cnts[i] = min(cnts[i-j*j]+1, cnts[i])
		}
	}
	return cnts[n]
}

func main() {
	start := time.Now().UnixNano()
	fmt.Println(numSquares(12))
	fmt.Println("运行时间", time.Now().UnixNano()-start)
}
