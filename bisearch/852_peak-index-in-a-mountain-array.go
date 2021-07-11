package main

import "fmt"

func main() {
	arr := []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}
	arr = []int{3, 5, 3, 2, 0}
	fmt.Println(arr, peakIndexInMountainArray(arr))
}

func peakIndexInMountainArray(arr []int) int {

	lw, hi := 0, len(arr)-1
	if hi < 2 {
		panic("not a MountainArray, count of the array less than 3.")
	}

	for lw <= hi {
		mid := lw + (hi-lw)/2
		if mid == 0 {
			mid += 1
		}
		if mid == len(arr)-1 {
			mid -= 1
		}
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
