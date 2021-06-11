package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	/*
		li := []int{7, 2, 1, 10}
		li = dedup(li)

		signs := getSigns(len(li)-1)
		for _, sign := range signs {
			if getResOf24(li, sign){
				fmt.Print(li, sign)
				break
			}
		}

	*/

	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input = strings.TrimSuffix(input, "\n")
	inputLi := strings.Split(input, " ")
	li := []int{}

	for _, in := range inputLi {
		n, err := strconv.Atoi(in)
		if err != nil {
			fmt.Println(err)
		}
		li = append(li, n)
	}
	// fmt.Println(inputLi)
	li = dedup(li)
	signs := getSigns(len(li) - 1)
	for _, sign := range signs {
		if getResOf24(li, sign) {
			fmt.Print(li, sign)
			break
		}
	}
	//fmt.Print(getResOf24(li, signs))
}

func getResOf24(li []int, signs []string) bool {
	nNum := len(li)
	nSigns := len(signs)
	if nNum-1 != nSigns {
		return false
	}
	stackLi := []int{}
	stackSi := []string{}
	for i := 0; i < nSigns; i++ {
		if i == 0 {
			stackLi = append(stackLi, li[i])
		}
		switch signs[i] {
		case "+":
			stackSi = append(stackSi, signs[i])
			stackLi = append(stackLi, li[i+1])
		case "-":
			stackSi = append(stackSi, signs[i])
			stackLi = append(stackLi, li[i+1])
		case "*":
			res := stackLi[len(stackLi)-1] * li[i+1]
			stackLi[len(stackLi)-1] = res
		case "/":
			res := stackLi[len(stackLi)-1] / li[i+1]
			stackLi[len(stackLi)-1] = res
		default:
		}
	}

	for i := 0; i < len(stackSi); i++ {
		switch stackSi[i] {
		case "+":
			res := stackLi[i] + stackLi[i+1]
			stackLi[i+1] = res
		case "-":
			res := stackLi[i] - stackLi[i+1]
			stackLi[i+1] = res
		}
	}
	return stackLi[len(stackLi)-1] == 24

}

// 获取全排列
func getSigns(n int) [][]string {
	if n == 0 {
		return nil
	}
	if n == 1 {
		return [][]string{[]string{"+"}, []string{"-"}, []string{"*"}, []string{"/"}}
	}

	resn := [][]string{[]string{"+"}, []string{"-"}, []string{"*"}, []string{"/"}}
	res := getSigns(n - 1)
	if res == nil {
		return resn
	}
	result := [][]string{}
	for _, r := range res {
		for _, rn := range resn {
			rn = append(rn, r...)
			result = append(result, rn)
		}
	}

	return result
}

func dedup(li []int) []int {

	maps := make(map[int]bool)
	for _, v := range li {
		if _, ok := maps[v]; !ok {
			maps[v] = true
		}
	}
	li = []int{}
	for k := range maps {
		li = append(li, k)
	}
	return li
}
