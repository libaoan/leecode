package main

import (
	"fmt"
)

func main() {

	haystack := "hello"
	needle := "ll"
	fmt.Print(strStr(haystack, needle))
}

// 基础算法, 算法复杂度 N*（字符出匹配算法的复杂度）
func strStr(haystack string, needle string) int {

	isMatch := func(subhay, needle string) bool {
		if len(subhay) < len(needle) {
			return false
		}
		return subhay[:len(needle)] == needle
	}

	if len(needle) == 0 {
		return 0
	}

	for i, _ := range haystack {
		if len(haystack[i:]) < len(needle) {
			return -1
		}
		if isMatch(haystack[i:], needle) {
			return i
		}
	}
	return -1
}

// todo: 其他优化算法待补充
func strStr2(haystack string, needle string) int {
	return 0
}
