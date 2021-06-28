package main

import (
	"fmt"
	"math"
)

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
	i := 1

	for {
		found := false
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if math.Abs(float64(rCenter-r))+math.Abs(float64(cCenter-c)) == float64(i) {
					found = true
					res = append(res, []int{r, c})
				}
			}
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
