package main

import "fmt"

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2))
}

func isHappy(n int) bool {

	maps := make(map[int]int)

	for {
		if _, ok := maps[n]; ok {
			return false
		}

		sum := 0
		for n >= 10 {
			w := n % 10
			sum += w * w
			n = n / 10
		}
		sum += n * n
		if sum == 1 {
			return true
		} else {
			maps[n] = sum
			n = sum
		}
	}

}
