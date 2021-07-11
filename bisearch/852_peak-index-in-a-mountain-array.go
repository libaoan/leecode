package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}
	arr = []int{3, 5, 3, 2, 0}
	fmt.Println(arr, peakIndexInMountainArray(arr))
}

func peakIndexInMountainArray(arr []int) int {

	if len(arr) < 3 {
		panic("not a MountainArray, count of the array less than 3.")
	}

	lw, hi := 1, len(arr)-2

	for lw <= hi {
		mid := lw + (hi-lw)/2
		vMid := arr[mid]
		vPre := arr[mid-1]
		vAft := arr[mid+1]
		switch {
		case vMid > vPre && vMid > vAft:
			return mid
		case vMid < vAft:
			lw = mid + 1
		case vMid < vPre:
			hi = mid - 1
		default:
			panic("not a MountainArray, array has two neigh equal element")
		}
	}
	panic("not a MountainArray")
}

// 采用sort标准库的二分查找法
func peakIndexInMountainArray2(arr []int) int {
	return sort.Search(len(arr)-1, func(i int) bool { return arr[i] > arr[i+1] })
}
