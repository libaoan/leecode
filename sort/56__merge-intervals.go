package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}
	//intervals = [][]int{[]int{1,4}, []int{0,2}, []int{3,5}}
	result := merge(intervals)
	for _, res := range result {
		fmt.Print(res, " ")
	}
}

// 基础算法 todo 调试不完全
func merge(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})

	result := [][]int{}
	i, j := 0, 1
	for j < len(intervals) {
		if intervals[i][1] < intervals[j][0] {
			for j < len(intervals) && intervals[j-1][1] >= intervals[j][0] {
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
