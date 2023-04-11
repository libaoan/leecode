package main

import (
	"fmt"
)

func main() {
	s := "PAYPALISHIRING"
	rs := convert(s, 3)
	fmt.Printf("the zigzag of `%s` is \n'%s'", s, rs)

	s = "PAYPALISHIRING"
	rs = convert(s, 4)
	fmt.Printf("the zigzag of `%s` is \n'%s'", s, rs)

}

// 调试通过
// todo leecode无法通过，报内存异常，待处理
func convert(s string, numRows int) string {
	res := format(s, numRows)
	s = ""
	//fmt.Println(res)
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			if res[i][j] == 0 {
				s += " "
			} else {
				s += fmt.Sprintf("%c", res[i][j])
			}

		}
		s += "\n"
	}
	s = ""
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(res); j++ {
			if res[j][i] == 0 {
				s += " "
			} else {
				s += fmt.Sprintf("%c", res[j][i])
			}
		}
		s += "\n"
	}
	return s
}

func format(s string, n int) [][]byte {
	res := make([][]byte, 0)
	i := 0
	for r := 0; r < n-1 && i < len(s); r++ {
		if r == 0 {
			if i+n < len(s[i:]) {
				res = append(res, []byte(s[i:i+n]))
				i = i + n
			} else {
				r := make([]byte, n)
				copy(r, []byte(s[i:]))
				res = append(res, r)
				return res
			}
		} else {
			row := make([]byte, n)
			row[n-1-r] = s[i]
			res = append(res, row)
			i++
		}
	}

	if i < len(s) {
		r := format(s[i:], n)
		for _, x := range r {
			res = append(res, x)
		}
		return res

	}
	return res
}
