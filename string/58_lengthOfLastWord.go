package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {

	cnt := 0
	s = strings.TrimSpace(s)
	for _, c := range s {
		if c == ' ' {
			cnt = 1
		} else {
			cnt += 1
		}
	}
	return cnt
}

func main() {
	var s string
	s = "Hello World"
	fmt.Println(lengthOfLastWord(s))
}
