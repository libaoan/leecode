package main

import (
	"fmt"
	"sort"
)

func longestPalindrome(s string) int {

	cnt := 0
	maps := make(map[int32]int)
	for _, c := range s {
		_, ok := maps[c]
		if ok {
			maps[c]++
		} else {
			maps[c] = 1
		}
	}

	type Kvs struct {
		Key   int32
		Value int
	}

	lists := make([]Kvs, 0)
	for k, v := range maps {
		if maps[k]%2 == 0 {
			cnt += maps[k]
		} else if maps[k] > 1 {
			cnt = cnt + maps[k] - 1
		}
		lists = append(lists, Kvs{Key: k, Value: v})
	}

	sort.Slice(lists, func(i, j int) bool {
		return lists[i].Value > lists[j].Value
	})

	fmt.Println(lists)

	for _, kvs := range lists {
		if kvs.Value%2 == 1 {
			return cnt + 1
		}
	}

	return cnt
}

// 优化1
func longestPalindrome2(s string) int {

	maps := map[int32]int{}
	for _, c := range s {
		maps[c]++
	}

	cnt := 0
	for _, v := range maps {
		if v&1 == 1 {
			cnt += v - 1
		} else {
			cnt += v
		}
	}

	if cnt < len(s) {
		cnt++
	}

	return cnt

}

func main() {

	s := "aaaaaccc"
	fmt.Println(longestPalindrome(s))
}
