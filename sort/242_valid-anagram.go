package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "anagram"
	t := "nagaram"

	s = "你好吗"
	t = "吗好你"
	fmt.Println(s, t, isAnagram(s, t))
}

// 排序法
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	bs := []rune(s)
	bt := []rune(t)

	sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
	sort.Slice(bt, func(i, j int) bool { return bt[i] < bt[j] })

	return string(bs) == string(bt)
}
