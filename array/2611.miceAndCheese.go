package main

import "sort"

// O(n), T=78%, M=32%
func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	max, n := 0, len(reward1)

	diff := make([]int, n)
	for i := 0; i < n; i++ {
		max += reward1[i]
		diff[i] = reward1[i] - reward2[i]
	}
	sort.Ints(diff)
	for i := 0; i < n-k; i++ {
		max -= diff[i]
	}

	return max
}
