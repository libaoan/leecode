package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange2(nums, target))
}

func searchRange(nums []int, target int) []int {

	ln := len(nums)
	lw, hi := 0, ln-1
	first, last := -1, -1

OUTER1:
	for lw <= hi {
		mid := lw + (hi-lw)>>1
		switch {
		case nums[mid] < target:
			lw = mid + 1
		case nums[mid] > target:
			hi = mid - 1
		default:
			if mid == 0 || nums[mid-1] != target {
				first = mid
				break OUTER1
			} else {
				hi = mid - 1
			}
		}
	}

	if first == -1 {
		return []int{first, last}
	}

	lw, hi = first, ln-1
OUTER2:
	for lw <= hi {
		mid := lw + (hi-lw)>>1
		switch {
		case nums[mid] < target:
			lw = mid + 1
		case nums[mid] > target:
			hi = mid - 1
		default:
			if mid == ln-1 || nums[mid+1] != target {
				last = mid
				break OUTER2
			} else {
				lw = mid + 1
			}
		}
	}

	return []int{first, last}

}

// 代码简化
func searchRange2(nums []int, target int) []int {
	ln := len(nums)
	lw, hi := 0, ln-1
	first, last := -1, -1

	for lw <= hi {
		mid := lw + (hi-lw)>>1
		if nums[mid] < target {
			lw = mid + 1
		} else {
			hi = mid - 1
		}
	}
	if lw < ln && nums[lw] == target {
		first = lw
	}

	if first != -1 {
		lw, hi = first, ln-1
		for lw <= hi {
			mid := lw + (hi-lw)>>1
			if nums[mid] > target {
				hi = mid - 1
			} else {
				lw = mid + 1
			}
		}
		if nums[hi] == target {
			last = hi
		}
	}

	return []int{first, last}
}
