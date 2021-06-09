package main

import "fmt"

func main() {

	n := 25
	fmt.Println(tribonacci(n))

}

var maps = make(map[int]int)

// 基础算法， 空间复杂度N
func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	if v, ok := maps[n]; ok {
		return v
	} else {
		maps[n] = tribonacci(n-3) + tribonacci(n-2) + tribonacci(n-1)
		return maps[n]
	}
}
