package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {

	maxlen, minlen := len(a), len(b)

	if maxlen < minlen {
		maxlen, minlen = minlen, maxlen
		a, b = b, a
	}

	var s string
	var i int = 1

	flag := uint8(0)

	for i <= minlen {
		sum := b[minlen-i] + a[maxlen-i] + flag - (48 * 2)
		if sum > 1 {
			flag = 1
			sum = sum % 2
		} else {
			flag = 0
		}
		s = strconv.Itoa(int(sum)) + s
		i++

	}

	for i <= maxlen {
		sum := byte(a[maxlen-i]) + flag - 48
		if sum > 1 {
			flag = 1
			sum = 0
		} else {
			flag = 0
		}

		s = strconv.Itoa(int(sum)) + s
		i++
	}

	if flag > 0 {
		s = strconv.Itoa(int(flag)) + s
	}

	return s
}

func main() {
	a, b := "1111", "1111"
	fmt.Println(addBinary(a, b))
}
