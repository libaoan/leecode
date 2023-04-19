package main

import (
	"fmt"
)

// todo: 部分通过，待调试
func findAllPeople(n int, meetings [][]int, firstPerson int) []int {

	ans := make([]int, n)
	ans[firstPerson] = firstPerson
	maybe := make(map[int][][]int, 0)
	for i := 0; i < len(meetings); i++ {
		if aa, ok := maybe[meetings[i][2]]; !ok {
			maybe[meetings[i][2]] = [][]int{meetings[i]}
		} else {
			for _, a := range aa {
				for _, x := range a {
					if meetings[i][0] == x {
						a = append(a, meetings[i][1])
						maybe[meetings[i][2]] = aa
						break
					}
					if meetings[i][1] == x {
						a = append(a, meetings[i][0])
						maybe[meetings[i][2]] = aa
						break
					}
				}
			}
		}

	}
	// fmt.Println("***", maybe)
	for i := 1; i < n; i++ {
		if aa, ok := maybe[i]; ok {
			for _, a := range aa {
				for _, x := range a {
					if ans[x] != 0 {
						for _, x := range a {
							ans[x] = x
						}
					}
				}
			}
		}
	}
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
	fmt.Println(findAllPeople(6, meetings, 1))
}
