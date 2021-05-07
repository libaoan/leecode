package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "()"
	s = "()[]{}"
	s = "(]"
	s = "([)]"
	s = "{[]}"
	res := isValid(s)
	res = isValid2(s)
	fmt.Printf("the str `%s` is valid parentheses %v", s, res)
}

func isValid(s string) bool {
	strInStack := "([{"
	strCompare := ")]}"

	var stack []string
	for _, c := range s {
		sc := string(c)
		switch {
		case strings.Contains(strInStack, sc):
			{
				stack = append(stack, sc)
			}
		case strings.Contains(strCompare, sc):
			{
				topStack := len(stack) - 1
				if topStack < 0 {
					return false
				} else {
					index := strings.Index(strCompare, sc)
					if stack != nil && (stack[topStack] != string(strInStack[index])) {
						return false
					} else {
						if topStack == 0 {
							stack = nil
						} else {
							stack = stack[0:topStack]
						}
					}
				}

			}
		default:
			{
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	} else {
		return true
	}
}

// 优化
func isValid2(s string) bool {
	length := len(s)
	if length%2 == 1 {
		return false
	}

	maps := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < length; i++ {
		if maps[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != maps[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
