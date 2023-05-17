package main

// O(n) 贪心算法 速度 95%, 内存 63%
func diStringMatch(s string) []int {
	lo, hi := 0, len(s)
	perm := make([]int, len(s)+1)
	for i, c := range s {
		if c == 'I' {
			perm[i] = lo
			lo++
		} else {
			perm[i] = hi
			hi--
		}
	}
	perm[len(s)] = lo
	return perm
}
