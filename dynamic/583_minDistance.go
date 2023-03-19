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

func main() {
	word1 := "sea"
	word2 := "eat"
	start := time.Now().UnixNano()
	fmt.Println(minDistance1(word1, word2))
	fmt.Println("运行时间", time.Now().UnixNano()-start)
}
