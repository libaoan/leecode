package main

import "fmt"

func main() {
	s := "PAYPALISHIRING"
	rs := convert(s, 3)
	fmt.Printf("the zigzag of `%s` is '%s'", s, rs)

}

// TODO: 未完成
func convert(s string, numRows int) string {
	length := len(s)
	if numRows < 1 {
		return ""
	}
	numCols := 0
	if numRows < 3 {
		numCols = 2 * (length / numRows)
	} else {
		numCols = 2 * (length / (2*numRows - 2))
	}
	res := make([][]string, numRows)
	for i := range res {
		res[i] = make([]string, numCols, '0')
	}
	rs := ""
	for i := range res {
		for j := range res[i] {
			rs += res[i][j]
		}
	}
	return rs
}
