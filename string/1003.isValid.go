package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aabbcc"
	s = "aababcbccabc"
	s = "abacbabcc"
	fmt.Println(isValid(s))
}

// 速度 68%， 内存70% todo： 需要用堆栈实现
func isValid(s string) bool {
	for len(s) > 0 {
		fmt.Println(string(s))

		if len(s)%3 != 0 {
			return false
		}
		if !strings.Contains(s, "abc") {
			return false
		}
		s = strings.ReplaceAll(s, "abc", "")

	}
	return true
}
