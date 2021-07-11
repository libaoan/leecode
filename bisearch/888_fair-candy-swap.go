package main

import "fmt"

func main() {
	as := []int{1, 2, 5}
	bs := []int{2, 4}
	fmt.Println(as, bs, fairCandySwap(as, bs))
}

// 二分查效率太差 哈希查找
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	SumX := 0
	mapX := make(map[int]bool)
	SumY := 0
	for _, v := range aliceSizes {
		SumX += v
		mapX[v] = true
	}

	for _, v := range bobSizes {
		SumY += v
	}

	delta := (SumX - SumY) / 2
	for i := range bobSizes {
		y := bobSizes[i]
		x := delta + y
		if _, ok := mapX[x]; ok {
			return []int{x, y}
		}
	}
	return nil
}
