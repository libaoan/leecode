package main

import "strconv"

// 速度 100%， 内存 80%
func alternateDigitSum(n int) int {

	sn := strconv.FormatUint(uint64(n), 10)
	sum, flag := 0, 1
	for _, i := range sn {
		n, _ := strconv.Atoi(string(i))
		sum += flag * n
		flag = -flag
	}
	return sum
}
