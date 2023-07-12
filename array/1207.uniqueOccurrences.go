package main

import "sort"

// T(O) = 100% S(O) = 15%
func uniqueOccurrences(arr []int) bool {
	maps := map[int]int{}
	for _, i := range arr {
		if _, ok := maps[i]; ok {
			maps[i]++
		} else {
			maps[i] = 1
		}
	}
	values := make([]int, 0, len(maps))
	for _, v := range maps {
		values = append(values, v)
	}
	sort.Ints(values)
	for i := 0; i < len(values); i++ {
		if i < len(values)-1 && values[i] == values[i+1] {
			return false
		}
	}

	return true
}
