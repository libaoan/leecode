package main

import "fmt"

func main() {
	var s string = "baca"
	ali, length := lengthOfLongestSubstring(s)
	fmt.Printf("lengthOfLongestSubstring of '%s' is '%s' with lenght %d", s, ali, length)
}

// 暴力法
func lengthOfLongestSubstring(s string) ([]byte, int) {
	var ali = []byte{}
	li := []byte(s)
	for i := 0; i < len(li); i++ {
		sli := []byte{byte(li[i])}
		for _, c := range li[i+1:] {
			if isByteInSlice(byte(c), sli) {
				if len(ali) < len(sli) {
					ali = sli
					break
				}
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
