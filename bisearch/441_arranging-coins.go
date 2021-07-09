package main

import "fmt"

func main() {
	fmt.Println(arrangeCoins2(1957747793))
}

// 基础算法， N*logN
func arrangeCoins(n int) int {

	fullArray := func(k int) int {
		if k < 1 {
			return 0
		}
		sum := 0
		for i := 1; i <= k; i++ {
			sum += i
		}
		return sum
	}

	lw, hi := 1, n
	for lw <= hi {
		mid := lw + (hi-lw)/2
		nMid := fullArray(mid)
		switch {
		case nMid > n:
			hi = mid - 1
		case nMid < n:
			lw = mid + 1
		default:
			return mid
		}
	}
	return hi
}

// 二分查找，优化全排列的计数算法，采用数学公式
func arrangeCoins2(n int) int {

	lw, hi := 1, n
	for lw <= hi {
		mid := lw + (hi-lw)/2
		nMid := mid * (mid + 1) / 2
		switch {
		case nMid > n:
			hi = mid - 1
		case nMid < n:
			lw = mid + 1
		default:
			return mid
		}
	}
	return hi

}
