package main

import (
	"fmt"
	"time"
)

// hash表， 速度72%, 内存72%
func containsDuplicate(nums []int) bool {
	var m = make(map[int]bool)

	for _, n := range nums {
		if _, ok := m[n]; ok {
			return true
		}
		m[n] = true
	}
	return false
}

func main() {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	nums = []int{1, 2, 3, 4}
	start := time.Now().UnixNano()
	fmt.Println(containsDuplicate(nums))
	fmt.Println("运行时间", time.Now().UnixNano()-start)
}
