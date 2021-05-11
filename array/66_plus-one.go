package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	nums = []int{4, 3, 2, 1}
	nums = []int{9, 9}
	res := plusOne2(nums)
	fmt.Printf("res + 1 : %v", res)
}

// 递归算法，最差情况：时间复杂度n， 空间复杂度n（当没位需要进位），且受制于n的长度
func plusOne(digits []int) []int {
	length := len(digits)

	digits[length-1] += 1
	if digits[length-1] < 10 {
		return digits
	} else {
		n := digits[length-1] - 10
		if length == 1 {
			digits = append([]int{1}, n)
		} else {
			digits = plusOne(digits[0 : length-1])
			digits = append(digits, n)
		}
		return digits
	}
}

// 动态规划算法
func plusOne2(digits []int) []int {
	length := len(digits)
	i := length - 1
	digits[i] += 1
	for ; i > 0; i-- {
		if digits[i] >= 10 {
			digits[i] -= 10
			digits[i-1] += 1
		}
	}

	if digits[i] >= 10 {
		digits[i] -= 10
		newDigits := make([]int, length+1)
		newDigits[0] = 1
		for j, n := range digits {
			newDigits[j+1] = n
		}
		digits = newDigits
	}

	return digits
}
