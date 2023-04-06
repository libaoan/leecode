package main

import "fmt"

// 1. p 不是点,且后面不是*， 直接匹配，不相同返回false，相同下一位
// 2. p 是点，且后面是星，直接匹配到p的倒数第星位
// 3. p 是点，s下一位
// todo : 通过率80%, 待完善
func isMatch(s string, p string) bool {
	sLen, pLen := len(s), len(p)

	if sLen == 0 && pLen == 0 {
		return true
	}
	if sLen == 0 || pLen == 0 {
		return false
	}

	if p[0] != '.' {
		// 0后面是*， i++， 直到0后面不是*
		// s跳到后面 i位匹配
		i := 1
		for i < pLen && p[i] == '*' {
			i++
		}
		// p已经结尾
		if i == pLen {
			// i== 1， 说明后面没有*
			if i == 1 {
				if s[0] == p[0] {
					if len(s) > 1 {
						return false
					} else {
						return true
					}
				} else {
					return false
				}
			}

			for j := 0; j < sLen; j++ {
				if s[j] != p[0] {
					return false
				}
			}
			return true
		} else {
			// i== 1， 说明后面没有*
			if i == 1 {
				if s[0] == p[0] {
					if len(s) > 1 {
						return isMatch(s[1:], p[1:])
					} else {
						return false
					}
				} else {
					return false
				}
			}

			j := 0
			for ; j < sLen; j++ {
				if s[j] != p[0] {
					return false
				}
			}
			if j == sLen {
				return false
			}
			return isMatch(s[j:], p[i:])
		}

	} else {
		// 0后面是*， i++， 直到0后面不是*
		// s跳到后面 i位匹配
		i := 1
		for i < pLen && p[i] == '*' {
			i++
		}
		// p已经结尾, 说明为 .** 模式，完全匹配
		if i == pLen {
			return true
		} else {
			if i >= sLen {
				return false
			}
			return isMatch(s[i:], p[i:])
		}
	}

}

func main() {
	s, p := "aa", "a"
	s, p = "aa", "a*"
	s, p = "ab", ".*"
	s, p = "aab", "c*a*b"
	fmt.Println(isMatch(s, p))
}
