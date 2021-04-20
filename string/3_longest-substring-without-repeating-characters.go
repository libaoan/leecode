package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "baca"
	s = "bbbbb"
	s = "pwwkew"
	maxS, length := lengthOfLongestSubstring(s)
	fmt.Printf("lengthOfLongestSubstring of '%s' is '%s' with lenght %d\n", s, maxS, length)

	SubS, lengSubS := lengthOfLongestSubstring2(s)
	fmt.Printf("lengthOfLongestSubstring of '%s' is '%s' with lenght %d\n", s, SubS, lengSubS)

	SubS, lengSubS = lengthOfLongestSubstring3(s)
	fmt.Printf("lengthOfLongestSubstring of '%s' is '%s' with lenght %d\n", s, SubS, lengSubS)
}

// 最初级的暴力法
func lengthOfLongestSubstring(s string) ([]byte, int) {
	var ali = []byte{}
	li := []byte(s)
	for i := 0; i < len(li); i++ {
		sli := []byte{byte(li[i])}
		for _, c := range li[i+1:] {
			if isByteInSlice(byte(c), sli) {
				if len(ali) < len(sli) {
					ali = sli
				}
				break
			}
			sli = append(sli, byte(c))
		}

	}
	return ali, len(ali)
}

func isByteInSlice(c byte, li []byte) bool {
	for _, x := range li {
		if byte(x) == c {
			return true
		}
	}
	return false
}

// 对暴力法优化
func lengthOfLongestSubstring2(s string) (string, int) {
	var maxS = ""
	length := len(s)

	for i := 0; i < length; i++ {
		subS := s[i : i+1]
		for j := i + 1; j < length; j++ {
			if strings.Contains(subS, s[j:j+1]) {
				if len(maxS) < len(subS) {
					maxS = subS
					break
				}
			} else {
				subS += s[j : j+1]
			}
		}
	}
	return maxS, len(maxS)
}

// 滑动窗口法
func lengthOfLongestSubstring3(s string) (string, int) {
	var maxS string = ""
	length := len(s)
	left := 0

	for i := 1; i < length; i++ {
		index := strings.Index(s[left:i], s[i:i+1])
		if index != -1 {
			if len(maxS) < len(s[left:i]) {
				maxS = s[left:i]
			}
			index = left + index
			left = index + 1
		}

	}

	return maxS, len(maxS)
}
