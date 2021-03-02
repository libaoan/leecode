package main

import "fmt"


func main(){
	fmt.Println(getBestPos1([]int{1,8,6,2,5,4,8,3,7}))
	fmt.Println(getBestPos2([]int{1,8,6,2,5,4,8,3,7}))
}


func getBestPos1(nums []int)(a1,a2,area int){
	n := len(nums)
	if n < 2{
		return 0, 0, area
	}
	a1, a2 = 0, 0
	area = 0
	for i, a := range nums{
		for j:=1; j<n; j++{
			w := j-i
			h := 0
			if a < nums[j]{
				h = a
			}else{
				h = nums[j]
			}
			if area < w*h{
				a1, a2 = i, j
				area = w*h
			}
		}
	}
	return a1, a2, area

}
// 优化1：双指针法
// 两条线段之间的面积受限与最短的线段，线段间距越长，面积越大
// 使用 2 个指针指向首部和尾部，将短指针向长指针方向移动，看能不能找到更长的线，使面积更大
// 依据：向长线方向每次移动 1 格，虽然宽度-1，但是(高度变高)*(宽度-1) >= 高度*宽度
//
func getBestPos2(nums []int)(a1,a2,area int){
	left, right := 0, len(nums)-1
	for left < right{
		w := right -left
		h := 0
		if nums[left] > nums[right]{
			h = nums[right]
		}else{
			h = nums[left]
		}

		if area < w*h{
			area = w*h
			a1, a2 = left, right
		}

		if nums[left] <= nums[right]{
			left++
		}else{
			right--
		}

	}
	return a1, a2,area
}
