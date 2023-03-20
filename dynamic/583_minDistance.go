package main

import (
	"fmt"
	"time"
)

// 回朔法， 运行超时
func minDistance1(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)

	if n1 == 0 && n2 == 0 {
		return 0
	}
	if n1 == 0 {
		return n2
	}
	if n2 == 0 {
		return n1
	}

	var cnt int
	if word1[0] == word2[0] {
		cnt = minDistance1(word1[1:], word2[1:])
	} else {
		cnt1 := minDistance1(word1[1:], word2) + 1
		cnt2 := minDistance1(word1, word2[1:]) + 1
		if cnt1 > cnt2 {
			cnt = cnt2
		} else {
			cnt = cnt1
		}
	}

	return cnt

}

// 动态规划 m[n1][n2] 代表 w[:n1] w[:n2] 的最小子序列
// 速度 88%， 内存22%
func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	if n1 == 0 {
		return n2
	}
	if n2 == 0 {
		return n1
	}
	m := make([][]int, n1+1)
	for i, _ := range m {
		m[i] = make([]int, n2+1)
	}
	for i := 0; i <= n2; i++ {
		m[0][i] = i
	}
	for i := 0; i <= n1; i++ {
		m[i][0] = i
	}
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			if word1[i] == word2[j] {
				m[i+1][j+1] = m[i][j]
			} else {
				m[i+1][j+1] = min(m[i][j+1], m[i+1][j]) + 1
			}
		}
	}
	return m[n1][n2]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	word1 := "a"
	word2 := "b"
	start := time.Now().UnixNano()
	fmt.Println(minDistance(word1, word2))
	fmt.Println("运行时间", time.Now().UnixNano()-start)
}
