package main

import "math"

func main() {
	println(titleToNumber("AB"))
	println(titleToNumber("ZY"))
}

// T(O) = 100%, S(O) = 95%
func titleToNumber(columnTitle string) int {
	bit, ans := len(columnTitle), 0
	for i := 0; i < len(columnTitle); i++ {
		bit--
		ans += int(columnTitle[i]-64) * int(math.Pow(float64(26), float64(bit)))
	}
	return ans
}
