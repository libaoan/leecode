package main

import "fmt"

func main() {
	ss := []string{"leetcode", "loveleetcode", "aabb"}
	for i := range ss {
		fmt.Println(ss[i], firstUniqChar(ss[i]))
		fmt.Println(ss[i], firstUniqChar2(ss[i]))
	}

}

// 线性查找， 时间复杂度 n**2
func firstUniqChar(s string) int {
	for i := range s {
		found := false
		for j := range s {
			if i != j && s[i] == s[j] {
				found = true
				break
			}
		}
		if !found {
			return i
		}
	}
	return -1
}

// 哈希查找, 空间复杂度N, 时间复杂度2N
func firstUniqChar2(s string) int {

	maps := make(map[byte]int)

	for i := range s {
		if _, ok := maps[s[i]]; ok {
			maps[s[i]]++
		} else {
			maps[s[i]] = 1
		}
	}

	for i := range s {
		if maps[s[i]] == 1 {
			return i
		}
	}
	return -1
}
