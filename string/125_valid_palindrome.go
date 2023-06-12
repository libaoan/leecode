package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome2("A man, a plan, a canal: Panama"))
	//fmt.Println(isPalindrome2("race a car"))
	fmt.Println(isPalindrome2("0P"))

}

// basic O(n) ， 速度 10% 内存6%
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	str := ""
	for _, r := range s {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
			str += string(r)
		}
	}

	head, tail := 0, len(str)-1
	for head < tail {
		if str[head] != str[tail] {
			return false
		}
		head++
		tail--
	}
	return true
}

// 优化 O(n) ， 速度 100% 内存95%
func isPalindrome2(s string) bool {

	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isalnum(s[left]) {
			left++
		}
		for left < right && !isalnum(s[right]) {
			right--
		}
		if left < right {
			if !isequal(s[left], s[right]) {
				return false
			}
			left++
			right--
		}
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func isequal(c1, c2 byte) bool {
	if c1 == c2 {
		return true
	}
	fmt.Println(c1, c2)
	if c1 > c2 {
		return c1-c2 == 32 && !(c2 >= '0' && c2 <= '9')
	} else {
		return c2-c1 == 32 && !(c2 >= '0' && c2 <= '9')
	}
}
