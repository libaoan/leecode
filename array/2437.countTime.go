package main

func main() {

}

// f(1), 速度 100%， 内存 99%
func countTime(time string) int {

	cntH, cntM := 1, 1
	if time[0] == '?' && time[1] == '?' {
		cntH = 24
	} else if time[0] == '?' && time[1] < '4' {
		cntH = 3
	} else if time[0] == '?' {
		cntH = 2
	} else if time[1] == '?' && time[0] == '2' {
		cntH = 4
	} else if time[1] == '?' {
		cntH = 10
	} else {
		cntH = 1
	}

	if time[3] == '?' && time[4] == '?' {
		cntM = 60
	} else if time[3] == '?' {
		cntM = 6
	} else if time[4] == '?' {
		cntM = 10
	} else {
		cntM = 1
	}

	return cntH * cntM

}
