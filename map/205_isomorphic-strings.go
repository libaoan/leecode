package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic3("egg", "add"))
	fmt.Println(isIsomorphic3("foo", "bar"))
	fmt.Println(isIsomorphic3("paper", "title"))
	fmt.Println(isIsomorphic3("bbbaaaba", "aaabbbba"))
	fmt.Println(isIsomorphic3("ab", "ac"))
	fmt.Println(isIsomorphic3("badc", "baba"))
}

// 基础算法，空间复杂度2N， 时间复杂度 2N
func isIsomorphic(s string, t string) bool {

	mapf := func(s, t string) bool {
		mapc := make(map[byte]byte)
		for i := range s {
			if _, ok := mapc[s[i]]; !ok {
				mapc[s[i]] = t[i]
			}
		}
		cs := ""
		for i := range s {
			c, _ := mapc[s[i]]
			cs += string(c)
		}

		if cs != t {
			return false
		}
		return true
	}

	return mapf(s, t) && mapf(t, s)

}

func isIsomorphic2(s string, t string) bool {
	s2t := map[byte]byte{}
	t2s := map[byte]byte{}
	for i := range s {
		x, y := s[i], t[i]
		if s2t[x] > 0 && s2t[x] != y || t2s[y] > 0 && t2s[y] != x {
			return false
		}
		s2t[x] = y
		t2s[y] = x
	}
	return true
}

// 贪心算法，优化， 空间复杂度
func isIsomorphic3(s string, t string) bool {
	maps := make(map[byte]byte)

	for i := range s {
		if c, ok := maps[s[i]]; !ok {
			// 不同的字符映射相同字符t[i]， 返回false
			for _, v := range maps {
				if v == t[i] {
					return false
				}
			}
			maps[s[i]] = t[i]
			// 相同的字符s[i]映射到不同的字符，返回false
		} else {
			if c != t[i] {
				return false
			}
		}
	}

	return true
}
