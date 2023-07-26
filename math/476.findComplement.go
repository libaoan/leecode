package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findComplement(5))
}

// T(O) = 100%  S(O) = 33%
func findComplement(num int) int {
	bins := strconv.FormatInt(int64(num), 2)
	s := ""
	for i := 0; i < len(bins); i++ {
		if bins[i] == '0' {
			s += "1"
		} else {
			s += "0"
		}
	}
	bins = s

	nnum, _ := strconv.ParseInt(bins, 2, len(bins))

	return int(nnum)
}
