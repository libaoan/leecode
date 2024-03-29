package main

import (
	"fmt"
	"sort"
)

// 贪心， nm
func findContentChildren(g []int, s []int) int {
	var cnt int

	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})

	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	for i := 0; i < len(g); i++ {
		for j := 0; j < len(s); j++ {
			if s[j] >= g[i] {
				cnt++
				s = append(s[0:j], s[j+1:]...)
				break
			}
		}
	}

	return cnt
}

// 贪心， n2
func findContentChildren2(g []int, s []int) int {
	var cnt int

	sort.Ints(s)
	sort.Ints(g)

	for i, j := 0, 0; i < len(g) && j < len(s); j++ {
		if g[i] <= s[j] {
			i++
			cnt++
		}
	}

	return cnt
}

func main() {
	g := []int{1, 1}
	s := []int{1, 1, 1}
	fmt.Println(findContentChildren2(g, s))
}
