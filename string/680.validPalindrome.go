package main

import "fmt"

func main() {
	s := "abc"
	fmt.Println(validPalindrome(s))
}

// O(n) 速度70%， 内存 50%
func validPalindrome(s string) bool {
	i := isPalindrome(s)
	if i == -1 {
		return true
	}
	n := len(s)
	println(i)

	if isPalindrome(s[i+1:n-i]) == -1 || isPalindrome(s[i:n-i-1]) == -1 {
		return true
	} else {
		return false
	}
}

func isPalindrome(s string) int {
	n := len(s)
	if n < 2 {
		return -1
	}
	i, j := 0, n-1
	for i <= j {
		if s[i] != s[j] {
			return i
		}
		i++
		j--
	}
	return -1
}
