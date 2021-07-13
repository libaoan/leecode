package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 60},
	}
	//nums = [][]int{[]int{1}, []int{3}}

	fmt.Println(searchMatrix(nums, 3))
	fmt.Println(searchMatrix(nums, 13))
	fmt.Println()
	fmt.Println(searchMatrix3(nums, 3))
	fmt.Println(searchMatrix3(nums, 13))

	fmt.Println()
	fmt.Println(searchMatrix2(nums, 3))
	fmt.Println(searchMatrix2(nums, 13))

	fmt.Println()
	fmt.Println(searchMatrix4(nums, 3))
	fmt.Println(searchMatrix4(nums, 13))
}

// 一次二分查找，把二维矩阵当作一维矩阵
func searchMatrix(matrix [][]int, target int) bool {
	rn := len(matrix)
	cn := len(matrix[0])
	tn := rn * cn

	lw, hi := 0, tn-1

	for lw <= hi {
		mid := lw + (hi-lw)>>2
		rmid := mid / cn
		cmid := mid % cn
		switch {
		case matrix[rmid][cmid] < target:
			lw = mid + 1
		case matrix[rmid][cmid] > target:
			hi = mid - 1
		default:
			return true
		}
	}
	return false
}

// 两次二分查找，列查找和行查找
func searchMatrix2(matrix [][]int, target int) bool {

	rn := len(matrix)
	cn := len(matrix[0])

	ir := -1

	lw, hi := 0, rn-1
	// 查找最后一位小于target的第0列
OUTER1:
	for lw <= hi {
		mid := lw + (hi-lw)>>2
		if matrix[mid][0] > target {
			hi = mid - 1
		} else {
			if mid == rn-1 || matrix[mid+1][0] > target {
				ir = mid
				break OUTER1
			} else {
				lw = mid + 1
			}
		}
	}

	if ir == -1 {
		return false
	}

	lw, hi = 0, cn-1
	// 查找目标行中的目标值
	for lw <= hi {
		mid := lw + (hi-lw)>>2
		switch {
		case matrix[ir][mid] > target:
			hi = mid - 1
		case matrix[ir][mid] < target:
			lw = mid + 1
		default:
			return true
		}
	}
	return false
}

// 用内置的sort包实现一次二分查找优化
func searchMatrix3(matrix [][]int, target int) bool {
	rn := len(matrix)
	cn := len(matrix[0])
	r := sort.Search(rn*cn, func(i int) bool {
		return matrix[i/cn][i%cn] >= target
	})
	return r < rn*cn && matrix[r/cn][r%cn] == target
}

// 用内置的sort包实现二次二分查找
func searchMatrix4(matrix [][]int, target int) bool {
	rn := len(matrix)
	cn := len(matrix[0])

	r := sort.Search(rn, func(i int) bool {
		return matrix[i][0] > target
	}) - 1

	if r < 0 {
		return false
	}

	index := sort.SearchInts(matrix[r], target)
	return index < cn && matrix[r][index] == target

}
