package main

import "fmt"

func main() {
	nums1 := "123"
	nums2 := "456"
	res := multiply(nums1, nums2)
	fmt.Printf("%d * %d = %d", nums1, nums2, res)
}

func multiply(num1 string, num2 string) string {
	sums := []int{}
	length1 := len(num1)
	length2 := len(num2)
	for i, _ := range num2 {
		jin := 0
		sum := 0
		cm := num2[length2-i-1] - 48
		for j, _ := range num1 {
			cn := num1[length1-j-1] - 48
			sSum := int(cn)*int(cm) + jin
			if sSum > 9 {
				jin = sSum / 10
				sSum = sSum % 10
			} else {
				jin = 0
			}
			wei := 1
			for k := 1; k <= j; k++ {
				wei = wei * 10
			}
			sum += wei * sSum
			if j == (length1-1) && jin != 0 {
				wei = wei * 10
				sum += wei * jin
			}
		}
		if sum > 0 {
			wei := 1
			for k := 1; k <= i; k++ {
				wei = wei * 10
			}
			sum = wei * sum
			sums = append(sums, sum)
		}
	}

	sum := 0
	for _, s := range sums {
		sum += s
	}
	str := []rune{}
	for _, c := range string(sum) {
		str = append(str, c)
	}
	return string(sum)
}
