package main

import "sort"

// todo 通过率 80%,审题错误
func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	max, n := 0, len(reward1)
	sort.Ints(reward1)
	sort.Ints(reward2)
	for i := n - 1; i >= n-k; i-- {
		max += reward1[i]
	}
	for i := n - 1; i >= k; i-- {
		max += reward2[i]
	}
	return max
}
