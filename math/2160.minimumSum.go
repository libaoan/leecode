package main

import "sort"

// T(O) = 100% S(O) = 85%
func minimumSum(num int) int {
	sn := []int{0, 0, 0, 0}
	for i := 0; i < 4 && num > 0; i++ {
		sn[i] = num % 10
		num = num / 10
	}
	sort.Ints(sn)
	return (sn[0]+sn[1])*10 + sn[2] + sn[3]
}
