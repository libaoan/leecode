package main

import "fmt"

func main() {
	res := allCellsDistOrder(2, 3, 1, 2)
	//res = allCellsDistOrder(2, 2, 0, 1)
	for _, r := range res {
		fmt.Println(r)
	}
}

// 基础算法，几何距离法
func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	res := [][]int{}

	res = append(res, []int{rCenter, cCenter})
	i, j := 0, 0

	for {
		found := false

		if rCenter-i >= 0 {
			found = true
			res = append(res, []int{rCenter - i, cCenter})
		}

		if cCenter-i >= 0 {
			found = true
			res = append(res, []int{rCenter, cCenter - i})
		}

		if rCenter+i < rows {
			found = true
			res = append(res, []int{rCenter + i, cCenter})
		}

		if cCenter+i < cols {
			found = true
			res = append(res, []int{rCenter, cCenter + i})
		}

		if rCenter-i >= 0 && cCenter-i >= 0 {
			found = true
			res = append(res, []int{rCenter - i, cCenter - i})
		}

		if rCenter-i >= 0 && cCenter+i < cols {
			found = true
			res = append(res, []int{rCenter - i, cCenter + i})
		}

		if rCenter+i < rows && cCenter-i >= 0 {
			found = true
			res = append(res, []int{rCenter + i, cCenter - i})
		}

		if rCenter+i < rows && cCenter+i < cols {
			found = true
			res = append(res, []int{rCenter + i, cCenter + i})
		}

		if found == false {
			return res
		} else {
			i++
		}

	}
}

// todo 桶排序
func allCellsDistOrder2(rows int, cols int, rCenter int, cCenter int) [][]int {
	res := [][]int{}
	return res
}
