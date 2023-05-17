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
	fmt.Println(largestNumber(nums))
}

// todo 通过率 90%
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
			return ss[i][k-1] >= ss[j][k]
		} else {
			return ss[i][k] >= ss[j][k-1]
		}
	})

	fmt.Println(ss)

	return strings.Join(ss, "")
}
