package main

import (
	"fmt"
	"sort"
)

// todo: 算法复杂度O(n4) 运行超时，通过率60%
func findAllPeople(n int, meetings [][]int, firstPerson int) []int {

	ans := make([]int, n)
	ans[firstPerson] = firstPerson
	maybe := make(map[int][][]int, 0)
	for i := 0; i < len(meetings); i++ {
		if aa, ok := maybe[meetings[i][2]]; !ok {
			maybe[meetings[i][2]] = [][]int{[]int{meetings[i][0], meetings[i][1]}}
		} else {
			match := false
			for j, a := range aa {
				for _, x := range a {
					if meetings[i][0] == x {
						match = true
						a = append(a, meetings[i][1])
						aa[j] = a
						maybe[meetings[i][2]] = aa
						break
					}
					if meetings[i][1] == x {
						match = true
						a = append(a, meetings[i][0])
						aa[j] = a
						maybe[meetings[i][2]] = aa
						break
					}
				}
			}
			if !match {
				aa = append(aa, []int{meetings[i][0], meetings[i][1]})
				maybe[meetings[i][2]] = aa
			}
		}

	}
	keys := make([]int, 0, len(maybe))
	for k := range maybe {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		aa := maybe[k]
		for _, a := range aa {
			for _, x := range a {
				if ans[x] == x {
					for _, x := range a {
						ans[x] = x
					}
				}
			}
		}
	}
	fmt.Println("***", maybe)

	an := []int{0}
	for i := 1; i < n; i++ {
		if ans[i] != 0 {
			an = append(an, ans[i])
		}
	}
	return an

}

func main() {
	meetings := [][]int{
		{3, 1, 3},
		{1, 2, 2},
		{0, 3, 3},
	}
	meetings = [][]int{
		{1, 2, 5},
		{2, 3, 8},
		{1, 5, 10},
	}
	meetings = [][]int{
		{3, 4, 2},
		{1, 2, 1},
		{2, 3, 1},
	}
	meetings = [][]int{
		{0, 2, 1},
		{1, 3, 1},
		{4, 5, 1},
	}
	fmt.Println(findAllPeople(6, meetings, 1))
}
