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
