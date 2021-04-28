package main

import "fmt"

func main() {
	ss := []string{"MCMXCIV", "LVIII", "IX", "IV", "III"}
	for _, s := range ss {
		fmt.Println(romanToInt2(s))
	}
}

func romanToInt(s string) int {
	maps := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}
	sum := 0
	for {
		length := len(s)
		switch {
		case length < 1:
			{
				return sum
			}
		case length == 1:
			{
				sum = sum + maps[s]
				return sum
			}
		case length > 1:
			{
				h1, h2 := string(s[0]), string(s[1])
				if maps[h1] >= maps[h2] {
					sum = sum + maps[h1]
					s = s[1:]
				} else {
					sum = sum + maps[h1+h2]
					s = s[2:]
				}
			}
		}
	}

}

// 优化，一步看2位， 前一位大为相加，前一位小为相减
func romanToInt2(s string) int {
	maps := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}
	sum := 0
	pre := maps[string(s[0])]
	for i := 1; i < len(s); i++ {
		n := maps[string(s[i])]
		if n > pre {
			sum = sum - pre
		} else {
			sum = sum + pre
		}
		pre = n
	}
	return sum + pre
}
