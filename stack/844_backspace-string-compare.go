package main

import "fmt"

func main() {
	s := "ab##"
	t := "c#d#"
	fmt.Printf("%s equals %s: %#v", s, t, backspaceCompare2(s, t))
}

// 基础算法，空间复杂度N，时间复杂度N，通过
func backspaceCompare(s string, t string) bool {

	filter := func(str string) string {
		stack := []byte{}
		for i := 0; i < len(str); i++ {
			if str[i] == '#' {
				if len(stack) != 0 {
					stack = stack[:len(stack)-1]
				}
			} else {
				stack = append(stack, str[i])
			}
		}
		return string(stack)
	}

	if filter(s) == filter(t) {
		return true
	}
	return false
}

// TODO：进阶算法，空间复杂度（1）  未通过
func backspaceCompare2(s string, t string) bool {

	finder := func(str string) (int, byte) {
		countDel := 0
		for i := len(str) - 1; i >= 0; i-- {
			switch {
			case str[i] != '#' && countDel > 0:
				countDel--
			case str[i] == '#':
				countDel++
			default:
				return i, str[i]
			}
		}
		return -1, '#'
	}
        # todo: there is problem，index should be len(s)
	index1, c1 := len(s)-1, s[len(s)-1]
	index2, c2 := len(t)-1, s[len(t)-1]

	for index1 > 0 && index2 > 0 {
		index1, c1 = finder(s[:index1])
		index2, c2 = finder(t[:index2])
		if index1 >= 0 && index2 >= 0 && c1 != c2 {
			return false
		}
	}

	if index1 > 0 {
		index1, c1 = finder(s[:index1])
		if index1 != -1 {
			return false
		}
		return true
	}

	if index2 > 0 {
		index2, c1 = finder(t[:index2])
		if index2 != -1 {
			return false
		}
		return true
	}
	return true
}
