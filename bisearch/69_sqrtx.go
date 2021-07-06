package main

import "math"

func main() {
	x := 2
	println(mySqrt(x))
}

// 二分查找法，近似值
func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	low := 0
	higth := x

	for low < higth {
		mid := low + (higth-low)/2
		switch {
		case mid*mid == x:
			return mid
		case mid*mid < x:
			low = mid + 1
		case mid*mid > x:
			higth = mid - 1
		}
	}
	return low

}

// 牛顿迭代法
func mySqrt2(x int) int {
	if x == 0 {
		return 0
	}

	C, x0 := float64(x), float64(x)
	for {
		xi := 0.5 * (x0 + C/x0)
		if math.Abs(x0-xi) < 1e-7 {
			break
		}
		x0 = xi
	}
	return int(x0)
}
