package main

import "fmt"

func main() {
	li := []int{7, 2, 1, 10}
	li = dedup(li)

	signs := []string{"*", "*", "+"}
	fmt.Print(getResOf24(li, signs))
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
func getSigns(n int) []string {
	signs := []string{"+", "-", "*", "/"}
	return signs
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
