package main

import (
	"fmt"
	"strconv"
)

func main() {
	ops := []string{"5", "2", "C", "D", "+"}
	//ops = []string{"5","-2","4","C","D","9","+","+"}
	ops = []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	score := calPoints2(ops)
	fmt.Println(score)
}

// 基础算法， 空间复杂度 N, 时间复杂度 2N
func calPoints(ops []string) int {
	stack := []int{}

	for _, x := range ops {
		switch {
		case x == "C":
			stack = stack[0 : len(stack)-1]
		case x == "D":
			score := 2 * stack[len(stack)-1]
			stack = append(stack, score)
		case x == "+":
			score := stack[len(stack)-1] + stack[len(stack)-2]
			stack = append(stack, score)
		default:
			x, _ := strconv.Atoi(x)
			stack = append(stack, x)
		}
	}

	score := 0
	for _, x := range stack {
		score += x
	}
	return score
}

// 优化算法， 空间复杂度 N, 时间复杂度 N
func calPoints2(ops []string) int {
	score := 0
	stack := []int{}

	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case "C":
			preScore := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			score -= preScore
		case "D":
			preScore := stack[len(stack)-1]
			tmpScore := 2 * preScore
			score += tmpScore
			stack = append(stack, tmpScore)
		case "+":
			preScore01 := stack[len(stack)-1]
			preScore02 := stack[len(stack)-2]
			tmpScore := preScore01 + preScore02
			score += tmpScore
			stack = append(stack, tmpScore)
		default:
			tmpScore, _ := strconv.Atoi(ops[i])
			score += tmpScore
			stack = append(stack, tmpScore)
		}
	}

	return score
}
