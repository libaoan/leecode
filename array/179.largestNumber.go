package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	nums := []int{1, 2, 3, 13}
	nums = []int{10, 2}
	nums = []int{3, 30, 34, 5, 9}
	nums = []int{34323, 3432}
	nums = []int{3333341, 3333}
	nums = []int{0, 0}
	nums = []int{12341, 123411234}
	fmt.Println(largestNumber(nums))
}

// todo 通过率 99%
func largestNumber(nums []int) string {

	n := len(nums)
	ss := make([]string, n)
	for i, c := range nums {
		ss[i] = fmt.Sprintf("%d", c)
	}

	sort.Slice(ss, func(i, j int) bool {
		li, lj := len(ss[i]), len(ss[j])
		if li == lj {
			return ss[i] >= ss[j]
		}
		k := 0
		for ; k < li && k < lj; k++ {
			if ss[i][k] > ss[j][k] {
				return true
			}
			if ss[i][k] < ss[j][k] {
				return false
			}
		}

		if k == li {
			for ; k < lj; k++ {
				for n := 0; n < li; n++ {
					if ss[i][n] > ss[j][k] {
						return true
					}
					if ss[i][n] < ss[j][k] {
						return false
					}
				}
			}
			return true
		} else {
			for ; k < li; k++ {
				for n := 0; n < lj; n++ {
					if ss[i][k] > ss[j][n] {
						return true
					}
					if ss[i][k] < ss[j][n] {
						return false
					}
				}
			}
			return true
		}
	})

	fmt.Println(ss)

	res := strings.TrimLeft(strings.Join(ss, ""), "0")
	if res == "" {
		return "0"
	} else {
		return res
	}
}
