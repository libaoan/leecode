package main

// O(n) 速度 62%， 内存93%
func lemonadeChange(bills []int) bool {
	maps := [2]int{0, 0}
	for _, b := range bills {
		switch {
		case b == 5:
			maps[0]++
		case b == 10:
			maps[0]--
			maps[1]++
			if maps[0] < 0 {
				return false
			}
		case b == 20:
			if maps[1] > 0 && maps[0] > 0 {
				maps[1]--
				maps[0]--
			} else if maps[0] > 2 {
				maps[0] = maps[0] - 3
			} else {
				return false
			}
		default:
			return false
		}
	}
	return true
}
