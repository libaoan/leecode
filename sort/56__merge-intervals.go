package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}
	//intervals = [][]int{[]int{1,4}, []int{0,2}, []int{3,5}}
	intervals = [][]int{
		[]int{1, 3}, []int{0, 2}, []int{2, 3}, []int{4, 6}, []int{4, 5}, []int{5, 5}, []int{0, 2}, []int{3, 3},
	}
	result := merge2(intervals)
	for _, res := range result {
		fmt.Print(res, " ")
	}
}

// 基础算法
func merge(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})

	result := [][]int{}
	i, j := 0, 1
	for j < len(intervals) {
		if intervals[i][1] < intervals[j][0] {
			for j < len(intervals) && intervals[j-1][1] >= intervals[j][0] {
				intervals[j][0] = intervals[j-1][0]
				if intervals[j-1][1] > intervals[j][1] {
					intervals[j][1] = intervals[j-1][1]
				}
				j++
			}
			if intervals[i][1] > intervals[j-1][1] {
				result = append(result, []int{intervals[i][0], intervals[i][1]})
			} else {
				result = append(result, []int{intervals[i][0], intervals[j-1][1]})
			}
			i = j
			j++
		} else {
			if intervals[j-1][1] >= intervals[j][0] {
				intervals[j][0] = intervals[j-1][0]
			}
			if intervals[j-1][1] > intervals[j][1] {
				intervals[j][1] = intervals[j-1][1]
			}
			j++
		}
	}

	if i >= len(intervals) {
		return result
	}
	if intervals[i][1] > intervals[j-1][1] {
		result = append(result, []int{intervals[i][0], intervals[i][1]})
	} else {
		result = append(result, []int{intervals[i][0], intervals[j-1][1]})
	}

	return result
}

// 别人的算法，我没想到，先实现一下
func merge2(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})

	result := [][]int{}
	ln := len(intervals)
	i, j := 0, 1
	t := intervals[0][1]
	max := intervals[ln-1][1]
	for _, x := range intervals {
		if max < x[1] {
			max = x[1]
		}
	}

	for t < max {
		if j < ln && t < intervals[j][0] {
			result = append(result, []int{intervals[i][0], t})
			i = j
			j++
			t = intervals[i][1]
		} else if j < ln {
			t = intervals[j][1]
			j++
		} else {
			result = append(result, []int{intervals[i][0], t})
		}
	}
	result = append(result, []int{intervals[i][0], t})
	return result
}
