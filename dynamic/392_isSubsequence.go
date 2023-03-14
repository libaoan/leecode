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

// 动态规划 速度100%， 内存60%
func isSubsequence(s string, t string) bool {

	m := len(s)
	if m == 0 {
		return true
	}
	n := len(t)
	cmp := -1
	for i := 0; i < m; i++ {
		j := cmp + 1
		for ; j < n; j++ {
			if s[i] == t[j] {
				cmp = j
				break
			}
		}
		// println(i,j)
		if j == n && i <= m-1 {
			return false
		}
	}

	return cmp != -1 && cmp < n
}

func main() {
	s := "acb"
	t := "ahbgdc"
	//s = "abc"
	//t = "aefbsc"
	start := time.Now().UnixNano()
	fmt.Println(isSubsequence(s, t))
	fmt.Println("运行时间", time.Now().UnixNano()-start)
}
