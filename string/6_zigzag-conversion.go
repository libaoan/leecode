package main

import "fmt"

func main() {
	s := "PAYPALISHIRING"
	rs := convert2(s, 3)
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

func convert2(s string, numRows int) string {
	res := format(s, numRows)
	ans := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		ans[i] = make([]byte, len(res))
	}

	s = ""
	fmt.Println(res)
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			ans[j][i] = res[i][j]
		}
	}
	fmt.Println(ans)
	for _, a := range ans {
		for _, c := range a {
			if c == 0 {
				s += " "
			} else {
				s += fmt.Sprintf("%c", c)
			}
		}
	}
	return s
}

func format(s string, n int) [][]byte {
	res := make([][]byte, 0)
	i := 0
	for r := 0; r < n && i < len(s); r++ {
		if r == 0 {
			if i+n < len(s[i:]) {
				res = append(res, []byte(s[i:i+n]))
				i = i + n
			} else {
				res = append(res, []byte(s[i:]))
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
			// fmt.Println("*", x)
			res = append(res, x)
		}
		//fmt.Println("**", res)
		return res

	}
	return res
}
