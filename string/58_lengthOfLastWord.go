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
			cnt = 0
		} else {
			cnt += 1
		}
	}

	return cnt
}

func lengthOfLastWord2(s string) int {

	cnt := 0
	index := len(s) - 1

	for s[index] == ' ' {
		index--
	}

	for index >= 0 && s[index] != ' ' {
		cnt++
		index--
		//fmt.Printf("%v, %v, %c\n", index, cnt, s[index])
	}

	return cnt
}

func main() {
	var s string
	s = "Hello World  "
	fmt.Println(lengthOfLastWord(s))
	fmt.Println(lengthOfLastWord2(s))
}
