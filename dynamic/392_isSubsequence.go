package main

import (
	"fmt"
	"time"
)

// 递归算法, 速度 14%， 内存 50%
func isSubsequence2(s string, t string) bool {

	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] != t[j] {
			if j == len(t)-1 {
				return false
			}
			if isSubsequence2(s[i:], t[j+1:]) {
				return true
			} else {
				return false
			}
		} else {
			i++
			j++
		}
	}
	if i == len(s) {
		return true
	}
	return false
}

// 动态规划
func isSubsequence(s string, t string) bool {
	n := len(t)
	cmp := -1
	for i := 0; i < len(s); i++ {
		for j := cmp + 1; j < n; j++ {
			if s[i] == t[j] {
				cmp = j
				break
			}
		}

	}
}

func main() {
	s := "axc"
	t := "ahbgdc"
	s = "abc"
	t = "aefbsc"
	start := time.Now().UnixNano()
	fmt.Println(isSubsequence2(s, t))
	fmt.Println("运行时间", time.Now().UnixNano()-start)
}
