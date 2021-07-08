package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	fmt.Println(firstBadVersion2(5))
}

func firstBadVersion(n int) int {

	lw, hi := 1, n

	for lw <= hi {
		index := lw + (hi-lw)/2
		switch {
		case isBadVersion(index):
			hi = index - 1
		default:
			lw = index + 1
		}
	}
	return lw
}

// 标准库自带的二分查找算法
func firstBadVersion2(n int) int {
	return sort.Search(n, isBadVersion)
}

func isBadVersion(version int) bool {
	rand.Seed(int64(version))
	b := []bool{true, false}
	return b[rand.Intn(2)]
}
