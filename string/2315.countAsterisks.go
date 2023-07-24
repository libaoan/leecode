package main

// T(O) = 100% S(O) = 70%
func countAsterisks(s string) int {
	cnt, flag := 0, true
	for i := 0; i < len(s); i++ {
		if s[i] == '|' {
			flag = !flag
		} else if s[i] == '*' && flag {
			cnt++
		}
	}
	return cnt
}
