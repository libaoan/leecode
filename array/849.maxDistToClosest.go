package main

import "fmt"

func main() {
	seats := []int{
		1, 0, 0, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0,
	}
	fmt.Println(maxDistToClosest(seats))
}

// todo P=90%
func maxDistToClosest(seats []int) int {
	ans, pos, n := -1, -1, len(seats)
	for i := 0; i < n; i++ {
		if seats[i] == 1 {
			seats[i] = 0
		} else if i == 0 {
			seats[i] = 1
		} else {
			seats[i] = seats[i-1] + 1
		}
	}

	for i := 0; i < n; i++ {
		if ans <= seats[i] {
			ans = seats[i]
			pos = i
		}
	}
	if ans == 1 || pos == ans-1 || pos == n-1 {
		return ans
	}
	return (ans + 1) / 2
}
