package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
}

// 哈希算法，双射的优化
func wordPattern(pattern string, s string) bool {
	sList := strings.Split(s, " ")
	maps := make(map[byte]string)

	if len(sList) != len(pattern) {
		return false
	}
	for i := range pattern {
		if s, ok := maps[pattern[i]]; !ok {
			for _, v := range maps {
				if v == sList[i] {
					return false
				}
			}
			maps[pattern[i]] = sList[i]
		} else {
			if s != sList[i] {
				return false
			}
		}
	}

	return true

}
